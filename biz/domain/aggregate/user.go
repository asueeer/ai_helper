package aggregate

import (
	"context"

	"ai_helper/biz/dal/db/repo"
	"ai_helper/biz/domain/entity"

	"github.com/jinzhu/gorm"
)

type UserAggregate struct {
	UserID      string             `json:"user_id"`
	User        entity.User        `json:"user"`
	UserAccount entity.UserAccount `json:"user_account"`
}

// Persist 持久化至数据库
func (userAggregate UserAggregate) Persist(ctx context.Context) error {
	userRepo := repo.NewUserRepo(ctx)
	err := userRepo.Do(ctx, func(sql *gorm.DB) error {
		if err := sql.Omit("id").Create(userAggregate.User.ToPo()).Error; err != nil {
			return err
		}
		if err := sql.Omit("id").Create(userAggregate.UserAccount.ToPo()).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}
