package main

import (
	"ai_helper/biz/handler"
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
			},
		),
	)

	r.GET("/api/ping", handler.Ping)

	// 用户中心
	r.GET("/api/profile/me", handler.ProfileMe)
	r.POST("/api/get_token", handler.RegisterVisitor)

	// 会话&消息
	r.POST("/api/im/create_conversation", handler.CreateConversation)
	r.POST("/api/im/send_message", handler.SendMessage, handler.SendMessageCallBack)
	r.POST("/api/im/load_conversation_detail", handler.LoadConversationDetail)
	r.POST("/api/im/load_conversations", handler.LoadConversations)
}
