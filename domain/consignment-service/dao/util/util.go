package util

import (
	"context"
	"demo1/domain/consignment-service/contextx"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// Define base model
type Model struct {
	ID        uint64 `gorm:"primaryKey;"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func GetDB(ctx context.Context, defDB *gorm.DB) *gorm.DB {
	trans, ok := contextx.FromTrans(ctx)
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
	return GetDB(ctx, defDB).Model(m)
}

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
