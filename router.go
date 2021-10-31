package main

import (
	"nearby/biz/handler"
	"nearby/biz/middleware"

	"github.com/gin-gonic/gin"
)

func register(r *gin.Engine) {
	r.Use(
		middleware.Cors(),
		middleware.Auth(
			middleware.JwtDefaultClient,
			map[string]bool{
				"/ping":      true,
				"/get_token": true,
			},
		),
	)

	r.GET("/ping", handler.Ping)

	// 用户中心
	r.GET("/profile/me", handler.ProfileMe)
	r.GET("/get_token", handler.RegisterVisitor)

	// 会话&消息
	r.GET("/im/create_conversation", handler.CreateConversation)
	r.GET("/im/send_message", handler.SendMessage)
	r.GET("/im/load_conversation_detail", handler.LoadConversationDetail)
	r.GET("/im/load_conversations", handler.LoadConversations)
}
