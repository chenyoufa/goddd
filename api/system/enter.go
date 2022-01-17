package system

import "thrgo/service"

type SystemApiGroup struct {
	BaseApi
	SystemApiApi
}

var (
	userService = service.ServiceGroupApp.SystemServiceGroup.UserService
	apiService  = service.ServiceGroupApp.SystemServiceGroup.ApiService
)
