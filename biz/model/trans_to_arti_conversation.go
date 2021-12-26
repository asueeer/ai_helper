package model

type TransToArtiConvRequest struct {
	ConvID string `json:"conv_id"`
}

type TransToArtiConvResponse struct {
	Meta Meta                   `json:"meta"`
	Data CreateConversationData `json:"data"`
}
