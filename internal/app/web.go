package app

import (
	"thrgoweb/internal/app/router"

	"github.com/gin-gonic/gin"
)

func InitGinEngine(r router.IRouter) *gin.Engine {
	app := gin.New()
	return app
}
