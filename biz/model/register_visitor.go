package model

type (
	RegisterVisitorRequest struct {
		UserID      string `form:"user_id" json:"user_id"`
		FingerPrint string `form:"finger_print" json:"finger_print"`
		VerifyCode  string `form:"user_id" json:"verify_code"`
	}

	RegisterVisitorData struct {
		Token          string `json:"token"`
		TokenExpiresAt int64  `json:"token_expires_at"` // token过期时间戳
	}

	RegisterVisitorResponse struct {
		Meta Meta                `json:"meta"`
		Data RegisterVisitorData `json:"data"`
	}
)
