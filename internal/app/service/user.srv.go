package service

import (
	"context"

	"github.com/google/wire"

	"thrgoweb/internal/app/dao"
)

var UserSet = wire.NewSet(wire.Struct(new(UserSrv), "*"))

type UserSrv struct {
	UserRepo *dao.UserRepo
}

func (a *UserSrv) UpdateStatus(ctx context.Context, id uint64, status int) error {

	return nil
}
