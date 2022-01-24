package api

import (
	"thrgoweb/internal/app/schema"
	"thrgoweb/internal/app/service"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

var UserSet = wire.NewSet(wire.Struct(new(UserAPI), "*"))

type UserAPI struct {
	UserSrc *service.UserSrv
}

func (a *UserAPI) Query(c *gin.Context){
	ctx:=c.Request.Context()
	var params schema.UserQueryParam
	if err:=ginx.
}