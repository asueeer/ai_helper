package vo

type Participant struct {
	UserID  int64  `json:"user_id"`  // 用户id
	HeadURL string `json:"head_url"` // 头像
}

type Conversation struct {
	ConvID       int64         `json:"conv_id"`      // 会话id
	Type         string        `json:"type"`         // 会话类型
	UnRead       int32         `json:"un_read"`      // 未读消息数
	LastMsg      Message       `json:"last_msg"`     // 最近一条消息
	Participants []Participant `json:"participants"` // 参与者
	ConvIcon     string        `json:"conv_icon"`    // 会话头像
	Timestamp    int64         `json:"timestamp"`    // 会话时间戳
}
