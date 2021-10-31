package model

type Response struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}

type Meta struct {
	Code int64  `json:"code"`
	Msg  string `json:"msg"`
}
