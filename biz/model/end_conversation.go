package model

type EndConversationRequest struct {
	ConvID string `form:"conv_id" json:"conv_id"`
}

type EndConversationResponse struct {
	Meta Meta `json:"meta"`
}
