package model

import "ai_helper/biz/model/vo"

type LoadConversationsRequest struct {
	Limit     int64  `form:"limit" json:"limit"`
	Cursor    string `form:"cursor" json:"cursor"`       // 相当于是offset
	Direction int32  `form:"direction" json:"direction"` // 拉取的方向; -1是由现在到过去; +1是由现在到未来
}

type LoadConversationData struct {
	Conversations []*vo.Conversation `json:"conversations"` // 会话列表
	NewCursor     string             `json:"new_cursor"`    // 下一次拉取所用的cursor
	HasMore       bool               `json:"has_more"`      // 是否还有更多
	Total         int64              `json:"total"`         // 会话总数
}

type LoadConversationsResponse struct {
	Meta Meta                 `json:"meta"`
	Data LoadConversationData `json:"data"`
}
