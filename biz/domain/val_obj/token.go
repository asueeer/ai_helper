package val_obj

type Token struct {
	Token          string `json:"token"`
	TokenExpiredAt int64  `json:"token_expired_at"`
}
