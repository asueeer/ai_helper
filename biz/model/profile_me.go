package model

type (
	ProfileMeRequest struct {
	}

	ProfileMeResponse struct {
		Meta Meta          `json:"meta"`
		Data ProfileMeData `json:"data"`
	}

	ProfileMeData struct {
		User User `json:"user"`
	}

	User struct {
		UserID   int64  `json:"user_id"`
		Nickname string `json:"nickname"`
		HeadURL  string `json:"head_url"`
	}
)
