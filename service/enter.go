package service

import "thrgo/service/system"

//各个功能块的总入口
type ServiceGroup struct {
	SystemServiceGroup system.SystemServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
