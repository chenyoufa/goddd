package system

import (
	"errors"
	"thrgo/global"
	"thrgo/model/system"

	"gorm.io/gorm"
)

type ApiService struct{}

var ApiServiceApp = new(ApiService)

func (apiService *ApiService) CreateApi(api system.SysApi) (err error) {
	if !errors.Is(
		global.GVA_DB.Where("path=? and method=?", api.Path, api.Method).First(&system.SysApi{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("存在相同api")
	}
	return global.GVA_DB.Create(&api).Error
}
