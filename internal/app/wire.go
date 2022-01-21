// +build wireinject

package app

import (
	"github.com/google/wire"
	"thrgoweb/internal/app/dao"
	"thrgoweb/internal/app/router"
	"thrgoweb/internal/app/service"
)

func BuildInjector() (*Injector, func(), error) {
	wire.Build(
		InitGormDB,
		dao.RepoSet,
		service.ServiceSet,
		InitGinEngine,
		router.RouterSet,
		InjectorSet,
	)
	return new(Injector), nil, nil
}
