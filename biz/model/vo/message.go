package vo

type Message struct {
	SenderID   int64      `json:"sender_id"`   // 发送方id
	ReceiverID int64      `json:"receiver_id"` // 接收方id
	Content    MsgContent `json:"content"`     // 消息内容
	Type       string     `json:"type"`        // 消息类型
	Status     string     `json:"status"`      // 消息状态
	Timestamp  int64      `json:"timestamp"`   // 消息时间戳
}

type MsgContent struct {
	Text     *string `json:"text,omitempty"`      // 纯文本
	RichText *string `json:"rich_text,omitempty"` // 富文本
	ImgURL   *string `json:"img_url,omitempty"`   // 图片链接
	AudioURL *string `json:"audio_url,omitempty"` // 语音链接
	VideoURL *string `json:"video_url,omitempty"` // 视频链接
}
