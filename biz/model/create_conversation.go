package model

type CreateConversationRequest struct {
	ReceiverID string `json:"receiver_id"` // 接收方id
	Type       string `json:"type"`        // 会话类型
}

type CreateConversationData struct {
	ConvID string `json:"conv_id"`
	IsNew  bool   `json:"is_new"`
	Status string `json:"status"`
}

type CreateConversationResponse struct {
	Meta Meta                   `json:"meta"`
	Data CreateConversationData `json:"data"`
}
