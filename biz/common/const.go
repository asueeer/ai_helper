package common

import "time"

// sso
const (
	UserProfile   = "user"
	AuthTokenName = "auth_token"
)

// 错误码
const (
	BizErrCode       = 50012 // 业务通用错误
	EvilViewErrCode  = 50013 // 非法查看不属于自己的资源
	LoginFailErrCode = 20010 // 用户未登陆
)

// 时间长度
const (
	Day     = time.Hour * 24
	Month   = Day * 30
	Year    = time.Hour * 24 * 365
	Century = Year * 100
)

// 会话业务相关
const (
	HelperConversationType = "helper" // 客服会话类型

	HelperConvStatusWaiting  = "roboting" // 客服会话状态-等待
	HelperConvStatusChatting = "chatting" // 聊天中
	HelperConvStatusEnd      = "end"      // 关闭

	HelperID = 435737 // 客服小助手的用户id

	ConvRoleCreator = "creator"   // 用户在会话中的角色-创建者
	ConvRoleVisitor = "visitor"   // 游客身份
	ConvRoleHelper  = "be_helper" // 后台客服
)
