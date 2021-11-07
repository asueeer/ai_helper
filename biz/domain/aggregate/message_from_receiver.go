package aggregate

import "ai_helper/biz/domain/entity"

// 接收者角度的消息
type MessageFromReceiver struct {
	MessageID   int64              `json:"message_id"`
	MessageFrom entity.MessageFrom `json:"message_from"`
	MessageTo   entity.MessageTo   `json:"message_to"`
}
