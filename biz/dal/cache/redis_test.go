package cache

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/spf13/cast"
)

func TestConnect(t *testing.T) {
	ctx := context.Background()
	for i := 0; i < 100; i++ {
		key := "foo" + cast.ToString(i)
		val := "bar" + cast.ToString(i)
		if Get(ctx, key).Val() != "" {
			t.Fatal(fmt.Sprintf("%v is not empty", key))
		}
		err := Set(ctx, key, val, time.Second*10)
		if err != nil {
			t.Fatal(fmt.Sprintf("set fail, err is %+v", err))
		}
		if Get(ctx, key).Val() != val {
			t.Fatal(fmt.Sprintf("Get(ctx, %s) != %s", key, val))
		}
	}
}

func TestGetFail(t *testing.T) {
	ctx := context.Background()
	key := "foo"
	err := Get(ctx, key).Err()
	t.Log(err)
}
