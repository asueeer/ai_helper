package model

type AcceptConversationRequest struct {
	ConvID string `form:"conv_id"  json:"conv_id"`
}

type AcceptConversationData struct {
	ConvID string `json:"conv_id"`
	Status string `json:"status"`
}

type AcceptConversationResponse struct {
	Data AcceptConversationData `json:"data"`
	Meta Meta                   `json:"meta"`
}
