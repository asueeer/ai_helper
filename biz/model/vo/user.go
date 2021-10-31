package vo

type User struct {
	UserID   string `json:"user_id"`  // 用户id, 用户唯一标示
	Nickname string `json:"nickname"` // 用户昵称
	HeadURL  string `json:"head_url"` // 用户头像链接
}
