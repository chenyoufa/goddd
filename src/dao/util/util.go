package util

import (
	"context"
	"goddd/src/contextx"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Model struct {
	ID        uint64 `gorm:"primaryKey;"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func GetDB(ctx context.Context, defDB *gorm.DB) *gorm.DB {
	trans, ok := contextx.FormTrans(ctx)
	if ok && !contextx.FromNoTrans(ctx) {
		db, ok := trans.(*gorm.DB)
		if ok {
			if contextx.FromTransLock(ctx) {
				db = db.Clauses(clause.Locking{Strength: "UPDATE"})
			}
			return db
		}
	}
	return defDB
}

func GetDBWithModel(ctx context.Context, defDB *gorm.DB, m interface{}) *gorm.DB {
	return GetDB(ctx, defDB)
}

type TransFunc func(context.Context) error

func ExecTrans(ctx context.Context, db *gorm.DB, fn TransFunc) error {
	transModel := &Trans{DB: db}
	return transModel.Exec(ctx, fn)
}

func ExecTransWithLock(ctx context.Context, db *gorm.DB, fn TransFunc) error {
	if !contextx.FromTransLock(ctx) {
		ctx = contextx.NewTransLock(ctx)
	}
	return ExecTrans(ctx, db, fn)
}

// func WrapPageQuery(ctx context.Context,db *gorm.DB,pp schema.pa)

// func FindPage(ctx context.Context,db *gorm.db,)

func FindOne(ctx context.Context, db *gorm.DB, out interface{}) (bool, error) {
	result := db.First(out)
	if err := result.Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func Check(ctx context.Context, db *gorm.DB) (bool, error) {
	var count int64
	result := db.Count(&count)
	if err := result.Error; err != nil {
		return false, err
	}
	return count > 0, nil
}

type OrderFieldFunc func(string) string

// func ParseOrder(items []*)
