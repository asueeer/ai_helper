package util

import (
	"encoding/json"
	"time"

	"github.com/spf13/cast"
)

func ToString(e interface{}) string {
	j, err := json.Marshal(e)
	if err != nil {
		return cast.ToString(e)
	}
	return string(j)
}

func ToTimeString(t time.Time) string {
	// todo 需要与产品对齐时间的格式
	return t.Format("2006-01-02 15:04:05")
}

func Micro2Sec(ms int64) (sec int64) {
	return ms / 1000
}

func Sec2Mirco(sec int64) (ms int64) {
	return sec * 1000
}

func NowUnixMicro() int64 {
	return time.Now().UnixMicro()
}
