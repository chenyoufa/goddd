package api

import (
	"goddd/src/ginx"

	"github.com/gin-gonic/gin"
)

func (a *UserAPI) Get(c *gin.Context) {
	ctx := c.Request.Context()
	item, err := a.UserSrv.Get(ctx, ginx.ParseParamID(c, "id"))
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	ginx.ResSuccess(c, item.CleanSecure())

}
