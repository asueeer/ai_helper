package service

import (
	"context"
	"testing"
)

func TestGenerateRandomNickname(t *testing.T) {
	var ss UserService
	ctx := context.Background()
	for i := 0; i < 10; i++ {
		t.Log(ss.GenerateRandomNickname(ctx))
	}
}
