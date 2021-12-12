package model

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Meta Meta      `json:"meta"`
	Data LoginData `json:"data"`
}

type LoginData struct {
	Token          string `json:"token"`
	TokenExpiresAt int64  `json:"token_expires_at"` // token过期时间戳
}
