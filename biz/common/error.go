package common

import (
	"fmt"

	"github.com/pkg/errors"
)

var NotLoginErr = &BizErr{
	Code: 0,
	Msg:  "",
	err:  nil,
}

type BizErr struct {
	Code int64
	Msg  string
	err  error
}

func (bizErr *BizErr) Error() string {
	return fmt.Sprintf(
		"code: %d, msg: %s, err: %s",
		bizErr.Code,
		bizErr.Msg,
		bizErr.err.Error(),
	)
}

func GetErrorCode(err error) int64 {
	if bizErr, ok := err.(*BizErr); !ok {
		return BizErrCode
	} else {
		return bizErr.Code
	}
}

func GetErrorMsg(err error) string {
	if bizErr, ok := err.(*BizErr); !ok {
		return err.Error()
	} else {
		return bizErr.Msg
	}
}

func NewBizErr(code int64, msg string, err error) *BizErr {
	if err == nil {
		return &BizErr{
			Code: code,
			Msg:  msg,
			err:  errors.New(""),
		}
	}
	return &BizErr{
		Code: code,
		Msg:  msg,
		err:  errors.Wrap(err, msg),
	}
}
