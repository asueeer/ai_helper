package entity

import "ai_helper/biz/dal/db/po"

type UserAccount struct {
	ID          int64  `json:"id"`           // 主键自增id,无业务意义
	UserID      int64  `json:"user_id"`      // 用户id, 用户唯一标示
	PhoneNumber string `json:"phone_number"` // 用户手机号
	WxOpenID    string `json:"wx_open_id"`   // 微信open_id, 微信侧唯一标识
}

func (account *UserAccount) ToPo() *po.UserAccount {
	return &po.UserAccount{
		ID:          account.ID,
		UserID:      account.UserID,
		PhoneNumber: account.PhoneNumber,
		WxOpenID:    account.WxOpenID,
	}
}
