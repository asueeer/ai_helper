package model

type CreateConversationRequest struct {
	ReceiverID string `json:"receiver_id"` // 接收方id
	Type       string `json:"type"`        // 会话类型
}

type CreateConversationData struct {
	ConvID string `json:"conv_id"`
}

type CreateConversationResponse struct {
	Meta Meta                   `json:"meta"`
	Data CreateConversationData `json:"data"`
}
