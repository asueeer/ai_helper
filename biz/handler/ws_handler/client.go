package ws_handler

import (
	"ai_helper/biz/common"
	"ai_helper/biz/domain/val_obj"
	"ai_helper/biz/model"
	"ai_helper/biz/service"
	"context"
	"encoding/json"
	"github.com/gorilla/websocket"
	"github.com/pkg/errors"
	"github.com/spf13/cast"
)

// Client is a middleman between the websocket connection and the hub.
type Client struct {
	// The websocket connection.
	conn *websocket.Conn

	user *val_obj.UserClaims

	uKey  string // 这个用户的长连接标识
	wsKey string // 这个连接的标识
}

func (c *Client) Register(ctx context.Context) error {
	// 存入hub
	c.user = common.GetUser(ctx)
	err := TheHub.Store(ctx, c)
	if err != nil {
		return common.NewBizErr(common.BizErrCode, err.Error(), err)
	}
	return nil
}

func (c *Client) Run(ctx context.Context) {
	ws := c.conn
	defer c.Close(ctx)
	for {
		//读取ws中的数据
		mt, message, err := ws.ReadMessage()
		if err != nil {
			break
		}
		// 处理收到的msg
		resp := handleMessage(ctx, message)
		//写入ws数据
		err = ws.WriteMessage(mt, resp)
		if err != nil {
			break
		}
	}
}

func (c *Client) Close(ctx context.Context) {
	c.conn.Close()
	TheHub.UnRegister(ctx, c)
}

func handleMessage(ctx context.Context, message []byte) []byte {
	var wsMsg model.WsMessage
	err := json.Unmarshal(message, &wsMsg)
	if err != nil {
		return []byte(errors.Wrap(err, "解析消息出错").Error())
	}
	ss := service.WsService{}
	switch wsMsg.Type {
	case common.WsHeartBeat:
		return []byte(cast.ToString(wsMsg.Msg))
	case common.WsLoadConv:
		return ss.WsLoadConv(ctx, wsMsg.Type, []byte(cast.ToString(wsMsg.Msg)))
	case common.WsLoadConvs:
		return ss.WsLoadConvs(ctx, wsMsg.Type, []byte(cast.ToString(wsMsg.Msg)))
	}
	return message
}
