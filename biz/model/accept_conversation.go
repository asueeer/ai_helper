package model

type AcceptConversationRequest struct {
	ConvID string `form:"conv_id"  json:"conv_id"`
}

type AcceptConversationResponse struct {
	Meta Meta `json:"meta"`
}
