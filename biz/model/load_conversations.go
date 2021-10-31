package model

import "nearby/biz/model/vo"

type LoadConversationsRequest struct {
	Limit  int64 `form:"limit" json:"limit"`
	Cursor int64 `form:"cursor" json:"cursor"` // 相当于是offset
}

type LoadConversationsResponse struct {
	Conversations []*vo.Conversation `json:"conversations"` // 会话列表
	NewCursor     int64              `json:"new_cursor"`    // 下一次拉取所用的cursor
	HasMore       bool               `json:"has_more"`      // 是否还有更多
	Total         int64              `json:"total"`         // 会话总数
}
