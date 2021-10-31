package model

type CreateConversationRequest struct {
	ReceiverID int64  `json:"receiver_id"` // 接收方id
	Type       string `json:"type"`        // 会话类型
}

type CreateConversationResponse struct {
	Meta   Meta  `json:"meta"`
	ConvID int64 `json:"conv_id"`
}
