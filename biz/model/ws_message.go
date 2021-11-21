package model

type WsMessage struct {
	Type int         `json:"type"`
	Msg  interface{} `json:"msg"`
}

type WsMessageResponse struct {
	Type int         `json:"type"`
	Msg  interface{} `json:"msg"`
}
