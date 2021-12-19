package main

import (
	"ai_helper/biz/handler"
	"ai_helper/biz/handler/ws_handler"
	"ai_helper/biz/middleware"

	"github.com/gin-gonic/gin"
)

func register(r *gin.Engine) {
	r.Use(
		middleware.Cors(),
		middleware.Auth(
			middleware.JwtDefaultClient,
			map[string]bool{
				"/api/ping":      true,
				"/api/get_token": true,
				"/api/login":     true,
			},
		),
	)

	r.GET("/api/ping", handler.Ping)

	// 用户中心
	r.GET("/api/profile/me", handler.ProfileMe)
	r.POST("/api/get_token", handler.RegisterVisitor)
	r.POST("/api/login", handler.Login)

	// 会话&消息
	r.POST("/api/im/create_conversation", handler.CreateConversation)
	r.POST("/api/im/send_message", handler.SendMessage, handler.SendMessageCallBack)
	r.POST("/api/im/load_conversation_detail", handler.LoadConversationDetail)
	r.POST("/api/im/load_conversations", handler.LoadConversations)

	// 长连接
	r.GET("/api/im/ws", ws_handler.WSHandler)

	// v2-会话工单需求
	r.POST("/api/im/accept_conversation", handler.AcceptConversation)
	r.POST("/api/im/end_conversation", handler.EndConversation)

	r.POST("/api/im/send_message_to_robot", handler.SendMessageToRobot)
}
