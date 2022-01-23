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
		UserID   string `json:"user_id"`
		Nickname string `json:"nickname"`
		HeadURL  string `json:"head_url"`
		IsHelper bool   `json:"is_helper"`
	}
)
