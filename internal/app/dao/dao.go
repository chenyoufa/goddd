package dao

import (
	"thrgoweb/internal/app/dao/user"

	"github.com/google/wire"
	"gorm.io/gorm"
)

//仓储统一注入
var RepoSet = wire.NewSet(
	user.UserSet,
)

//仓储统一入库
type (
	UserRepo = user.UserRepo
)

// Auto migration for given models
func AutoMigrate(db *gorm.DB) error {
	// if dbType := config.C.Gorm.DBType; strings.ToLower(dbType) == "mysql" {
	// 	db = db.Set("gorm:table_options", "ENGINE=InnoDB")
	// }

	return db.AutoMigrate(
		// new(menu.MenuActionResource),
		// new(menu.MenuAction),
		// new(menu.Menu),
		// new(role.RoleMenu),
		// new(role.Role),
		// new(user.UserRole),
		new(user.User),
	) // end
}
