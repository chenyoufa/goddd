package system

import (
	"thrgo/global"
	"thrgo/model/common/request"
	"thrgo/model/common/response"

	// us "thrgo/service/system"
	"thrgo/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// func (b *BaseApi) Login(c *gin.Context) {
// 	var l systemReq.Login
// 	_ = c.ShouldBindJSON(&l)
// 	if err := utils.Verify(l, utils.LoginVerify); err != nil {
// 		response.FailWithMessage(err.Error(), c)
// 	}
// 	if store.Verify(l.CaptchaId, l.Captcha, true) {
// 		u := &system.SysUser{Username: l.Username, Password: l.Password}
// 		if err, user := userService.Login(u); err != nil {
// 			global.GVA_LOG.Error("登陆失败! 用户名不存在或者密码错误!", zap.Error(err))
// 			response.FailWithMessage("用户名不存在或者密码错误", c)
// 		} else {
// 			// b.tokenNext(c, *&user)
// 		}
// 	} else {
// 		response.FailWithMessage("验证码错误", c)
// 	}
// }

func (b *BaseApi) GetUserList(c *gin.Context) {
	var pageinfo request.PageInfo
	_ = c.ShouldBindJSON(&pageinfo)
	if err := utils.Verify(pageinfo, utils.PageInfoVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
	}
	// var uc = us.UserService{}
	// uc.GetUserInfoList(pageinfo)

	if list, total, err := userService.GetUserInfoList(pageinfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageinfo.Page,
			PageSize: pageinfo.PageSize,
		}, "获取成功", c)
	}
}
