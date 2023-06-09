package ws_handler

import (
	"ai_helper/biz/common"
	"ai_helper/biz/config"
	"ai_helper/biz/dal/cache"
	"ai_helper/biz/model"
	"context"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"log"
	"sync"
	"time"
)

type UserID int64

var maxConnCnt = 1500

var TheHub *Hub

func init() {
	TheHub = &Hub{
		manager: sync.Map{},
	}
}

type Hub struct {
	manager sync.Map
	cnt     int
}

func (h *Hub) Store(ctx context.Context, cli *Client) error {
	if h.cnt > maxConnCnt {
		return errors.New("已经超过了最大连接数")
	}
	log.Printf("user_%+v上线啦\n", cli.user.UserID)
	wsKey := fmt.Sprintf("user_id: %d, rand: %d", cli.user.UserID, config.GenerateIDInt64())
	uKey := fmt.Sprintf("ws_%d", cli.user.UserID)
	if cli.user.IsHelper {
		// 如果用户是客服, 把他在客服里注册上
		log.Printf("客服在登陆")
		wsKey = fmt.Sprintf("user_id: %d, rand: %d", common.HelperID, config.GenerateIDInt64())
		uKey = fmt.Sprintf("ws_%d", common.HelperID)
	}
	// 1. 在hub里注册client
	h.manager.Store(wsKey, cli)
	h.cnt++

	// 2. 将ticket存入redis, 之后根据ticket去hub里找client
	cache.SAdd(ctx, uKey, wsKey)
	log.Printf("uKey: %+v\n", uKey)
	cache.ExpireAt(ctx, uKey, time.Now().Add(time.Hour*24*365))
	cli.uKey = uKey
	cli.wsKey = wsKey
	return nil
}

func (h *Hub) UnRegister(ctx context.Context, cli *Client) {
	h.manager.Delete(cli.wsKey)
	h.cnt--
	cache.SRemove(ctx, cli.uKey, cli.wsKey)
}

func (h *Hub) Load(key string) *Client {
	cli, ok := h.manager.Load(key)
	if !ok {
		return nil
	}
	return cli.(*Client)
}

func (h *Hub) GetWsKeys(ctx context.Context, uKey string) []string {
	ss, err := cache.SMembers(ctx, uKey).Result()
	if err != nil {
		log.Printf("err: %+v", err)
		return nil
	}
	return ss
}

func (h *Hub) BatchSendMsgs(ctx context.Context, receiverID int64, msgNotify model.WsMessageResponse) {
	// 后面需要不止是通过id发送，还应该给固定的角色去发送消息...
	log.Printf("给长连接发送消息, receiverID is %+v", receiverID)
	wsKeys := h.GetWsKeys(ctx, fmt.Sprintf("ws_%d", receiverID))
	if len(wsKeys) == 0 {
		log.Printf("receiverID_%+v不在线\n", receiverID)
		return
	}
	for i := range wsKeys {
		log.Printf("wsKey is %+v", wsKeys[i])
		cli := h.Load(wsKeys[i])
		if cli == nil {
			continue
		}
		j, err := json.Marshal(msgNotify)
		if err != nil {
			log.Printf("SendMessageCallBack fail, err: %+v", err)
		}
		err = cli.WriteMessage(ctx, j)
		if err != nil {
			log.Printf("cli.WriteMessage fail, err: %+v", err)
		}
	}
}
