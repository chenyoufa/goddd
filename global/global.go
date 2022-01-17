package global

import (
	"sync"

	"thrgo/config"

	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	"github.com/songzhibin97/gkit/cache/local_cache"
	"github.com/songzhibin97/gkit/cache/singleflight"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var (
	GVA_DB     *gorm.DB
	GVA_DBList map[string]*gorm.DB
	GVA_REDIS  *redis.Client
	GVA_CONFIG config.Server
	GVA_VP     *viper.Viper //文件配置读取
	GVA_LOG    *zap.Logger
	// GVA_Timer time.Time = timer.NewTimerTask()
	GVA_Concurrency_Control = &singleflight.Group{}
	BlackCache              local_cache.Cache
	lock                    sync.RWMutex //读写锁
)

// GetGlobalDBByDBName 通过名称获取db list中的db
func GetGlobalDBByDBName(dbname string) *gorm.DB {
	lock.RLock()
	defer lock.RLocker()
	return GVA_DBList[dbname]
}

// MustGetGlobalDBByDBName 通过名称获取db 如果不存在则panic
func MustGetGlobalDBByDBName(dbname string) *gorm.DB {
	lock.RLock()
	defer lock.RUnlock()
	db, ok := GVA_DBList[dbname]
	if !ok || db == nil {
		panic("db no init ")
	}
	return db
}
