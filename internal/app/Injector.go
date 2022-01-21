package app

import (
	"thrgoweb/internal/app/service"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

var InjectorSet = wire.NewSet(wire.Struct(new(Injector), "*"))

type Injector struct {
	Engine *gin.Engine
	// Auth           auth.Auther
	// CasbinEnforcer *casbin.SyncedEnforcer
	UserSrv *service.UserSrv
}
