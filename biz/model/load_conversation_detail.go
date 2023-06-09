package model

import "ai_helper/biz/model/vo"

type LoadConversationDetailRequest struct {
	ConvID    string `form:"conv_id" json:"conv_id" binding:"required"` // 会话id
	Cursor    string `form:"cursor" json:"cursor"`                      // 客户端本地消息的存储位置(时间戳)，用于从服务端获取该位置之后的消息
	Limit     int64  `form:"limit" json:"limit"`                        // 默认50条
	Role      string `form:"role" json:"role"`                          // 用户身份; "visitor" 游客; "be_helper": 后台客服
	Direction int32  `form:"direction" json:"direction"`                // 拉取方向; 枚举值 +1: 由现在到未来; -1: 由现在到过去 （默认为-1）
}

type WsMessageLoadConvDetailRequest struct {
	Type int                           `json:"type"`
	Msg  LoadConversationDetailRequest `json:"msg"`
}

type LoadConversationDetailData struct {
	Messages  []*vo.Message `json:"messages"`   // 消息列表
	HasMore   bool          `json:"has_more"`   // 是否包含更多会话
	NewCursor string        `json:"new_cursor"` // 下一次拉取前, 需要传给后端的时间戳
}

type LoadConversationDetailResponse struct {
	Meta Meta `json:"meta"`

	Data LoadConversationDetailData `json:"data"`
}
