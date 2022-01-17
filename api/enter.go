package api

import (
	"thrgo/api/system"
)

type ApiGroup struct {
	SystemApiGroup system.SystemApiGroup
	// ExampleApiGroup  example.ApiGroup
	// AutoCodeApiGroup autocode.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
