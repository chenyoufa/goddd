package util

import (
	"context"
	"goddd/src/contextx"

	"github.com/google/wire"

	"gorm.io/gorm"
)

type Trans struct {
	DB *gorm.DB
}

var TransSet = wire.NewSet(wire.Struct(new(Trans), "*"))

func (a *Trans) Exec(ctx context.Context, fn func(context.Context) error) error {
	if _, ok := contextx.FormTrans(ctx); ok {
		return fn(ctx)
	}
	return a.DB.Transaction(func(db *gorm.DB) error {
		return fn(contextx.NewTrans(ctx, db))
	})
}
