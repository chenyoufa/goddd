package user

import (
	"context"
	"thrgoweb/internal/app/schema"
	"thrgoweb/internal/app/util"
	"thrgoweb/pkg/util/structure"

	"gorm.io/gorm"
)

//1每个request 一个 gotinue ，超时停止数据库等资源访问， token cookie 各个 reques 共享
//2  封装事物和其他gorm 操作
func GetUserDB(ctx context.Context, defDB *gorm.DB) *gorm.DB {
	return util.GetDBWithModel(ctx, defDB, new(User))
}

type User struct {
	util.Model
	UserName string  `gorm:"size:64;uniqueindex;default:'';not null;"`
	RealName string  `gorm:"size:64;index;default:'';"`
	Password string  `gorm:"size:40;default:'';"`
	Email    *string `gorm:"size:255;"`
	Phone    *string `gorm:"size:20"`
	Status   int     `gorm:"index;default:0"`
	Creator  uint64  `gorm:""`
}

func (a User) ToSchemaUser() *schema.User {
	item := new(schema.User)
	structure.Copy(a, item)
	return item
}

type Users []*User

func (a Users) ToSchemaUsers() []*schema.User {
	list := make([]*schema.User, len(a))
	for i, item := range a {
		list[i] = item.ToSchemaUser()
	}
	return list
}
