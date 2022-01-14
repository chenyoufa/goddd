package user

import (
	"context"
	"goddd/src/dao/util"
	"goddd/src/schema"

	"goddd/pkg/errors"

	"github.com/google/wire"
	"gorm.io/gorm"
)

type UserRepo struct {
	DB *gorm.DB
}

var UserSet = wire.NewSet(wire.Struct(new(UserRepo), "*"))

func (a *UserRepo) getQueryOption(opts ...schema.UserQueryOptions) schema.UserQueryOptions {
	var opt schema.UserQueryOptions
	if len(opts) > 0 {
		opt = opts[0]
	}
	return opt
}

func (a *UserRepo) Get(ctx context.Context, id uint64, opts ...schema.UserQueryOptions) (*schema.User, error) {
	var item User
	ok, err := util.FindOne(ctx, GetUserDB(ctx, a.DB).Where("id=?", id), &item)
	if err != nil {
		return nil, errors.WithStack(err)
	} else if !ok {
		return nil, nil
	}
	return item.ToSchemaUser(), nil
}
