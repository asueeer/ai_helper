package vo

type Message struct {
	MessageID  string     `json:"message_id"`  // 消息id
	ConvID     string     `json:"conv_id"`     // 会话id
	SenderID   string     `json:"sender_id"`   // 发送方id
	ReceiverID string     `json:"receiver_id"` // 接收方id
	Content    MsgContent `json:"content"`     // 消息内容
	Type       string     `json:"type"`        // 消息类型
	Status     string     `json:"status"`      // 消息状态
	SeqID      int64      `json:"timestamp"`   // 用于保序的序列号(历史原因，现在json字段是timestamp之后通知前端进行改动)
	Role       string     `json:"role"`        // 消息发送者的身份, 游客:"visitor"; 客服:"helper"
}

type MsgContent struct {
	Text     *string `json:"text,omitempty"`      // 纯文本
	RichText *string `json:"rich_text,omitempty"` // 富文本
	ImgURL   *string `json:"img_url,omitempty"`   // 图片链接
	AudioURL *string `json:"audio_url,omitempty"` // 语音链接
	VideoURL *string `json:"video_url,omitempty"` // 视频链接
	Link     *string `json:"link,omitempty"`      // 业务名对应的网站
	End      *string `json:"end"`                 // 诊断结束标志
}
