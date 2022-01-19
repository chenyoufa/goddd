package user

import (
	"context"
	"thrgoweb/internal/app/util"

	"gorm.io/gorm"
)

func GetUserDB(ctx context.Context, defDB *gorm.DB) *gorm.DB {
	return util.GetDBWithModel(ctx, defDB, new(User))
}

type User struct {
	util.Model
	UserName string  `gorm:"size:64;uniqueindex;default:'';not null;"`
	RealName string  `gorm:"size:64;index;default:'';"`
	Password string  `gorm:"size:255;"`
	Email    *string `gorm:"size:20;"`
	Phone    *string `gorm:"size"`
	Status   int     ``
	Creator  uint64  ``
}
