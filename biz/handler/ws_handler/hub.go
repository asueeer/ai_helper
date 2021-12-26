package ws_handler

import (
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

var maxConnCnt = 150

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

	wsKey := fmt.Sprintf("user_id: %d, rand: %d", cli.user.UserID, config.GenerateIDInt64())
	uKey := fmt.Sprintf("ws_%d", cli.user.UserID)
	// 1. 在hub里注册client
	h.manager.Store(wsKey, cli)
	h.cnt++

	// 2. 将ticket存入redis, 之后根据ticket去hub里找client
	cache.SAdd(ctx, uKey, wsKey)
	cache.ExpireAt(ctx, uKey, time.Now().Add(time.Hour*10))
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
	wsKeys := h.GetWsKeys(ctx, fmt.Sprintf("ws_%d", receiverID))
	if len(wsKeys) == 0 {
		return
	}
	for i := range wsKeys {
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
