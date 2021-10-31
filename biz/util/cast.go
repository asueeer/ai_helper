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
