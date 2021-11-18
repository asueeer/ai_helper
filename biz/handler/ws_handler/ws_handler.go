package ws_handler

import (
	"ai_helper/biz/common"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// WSHandler 升级长连接
func WSHandler(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Printf("err: %+v", err)
		common.WriteError(c, err)
		return
	}
	client := Client{
		conn: conn,
	}
	// 1. 在hub中注册
	err = client.Register(c)
	if err != nil {
		common.WriteError(c, err)
		return
	}

	// 2. 收发消息
	client.Run(c)
}
