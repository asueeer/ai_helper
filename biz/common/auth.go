package common

import (
	"context"
	"log"
	"nearby/biz/domain/val_obj"
)

func GetUser(ctx context.Context) *val_obj.UserClaims {
	v := ctx.Value(UserProfile)
	log.Printf("ctx.Value('user') is: %+v", v)
	if v == nil {
		return nil
	}
	user, ok := v.(*val_obj.UserClaims)
	if !ok || user == nil {
		return nil
	}
	return user
}
