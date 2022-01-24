package ginx

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"thrgoweb/internal/app/schema"
	"thrgoweb/pkg/errors"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

const (
	prefix     = "gin-admin"
	ReqBodyKey = prefix + "/req-body"
	ResBodyKey = prefix + "/res-body"
)

func GetToken(c *gin.Context) string {
	var token string
	auth := c.GetHeader("Authorization")
	prefix := "Bearer"
	if auth != "" && strings.HasPrefix(auth, prefix) {
		token = auth[len(prefix):]
	}
	return token
}

func GetBodyData(c *gin.Context) []byte {
	if v, ok := c.Get(ReqBodyKey); ok {
		if b, ok := v.([]byte); ok {
			return b
		}
	}
	return nil
}
func ParseParamID(c *gin.Context, key string) uint64 {
	val := c.Param(key)
	id, err := strconv.ParseUint(val, 10, 64)
	if err != nil {
		return 0
	}
	return id
}
func ParseJSON(c *gin.Context, obj interface{}) error {
	if err := c.ShouldBindJSON(obj); err != nil {
		return errors.Wrap400Response(err, fmt.Sprintf("Parse request json failed: %s", err.Error()))
	}
	return nil
}

func ParseQuery(c *gin.Context, obj interface{}) error {
	if err := c.ShouldBindQuery(obj); err != nil {
		return errors.Wrap400Response(err, fmt.Sprintf("Parse request query failed:%s", err.Error()))
	}
	return nil
}

func ParseForm(c *gin.Context, obj interface{}) error {
	if err := c.ShouldBindJSON(obj, binding.Form); err != nil {
		return errors.Wrap400Response(err, fmt.Sprintf("Parse request form failed: %s", err.Error()))
	}
	return nil
}

func ResOk(c *gin.Context) {

}

func ResList(c **gin.Context, v interface{}) {

}

func ResPage(c *gin.Context, v interface{}, pr *schema.PaginationResult) {
	list := schema.ListResult{
		List:       v,
		Pagination: pr,
	}

}

func ResSuccess(c *gin.Context, v interface{}) {

}

func ResJSON(c *gin.Context, status int, v interface{}) {
	buf, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}
	c.Set(ResBodyKey, buf)
	c.Data(status, "application/json:charset=utf-8", buf)
	c.Abort()
}

func ResError(c *gin.Context, err error, status ...int) {
	ctx := c.Request.Context()
	var res *errors.ResponseError
	if err != nil {
		if e, ok := err.(*errors.ResponseError); ok {
			res = e
		} else {
			res = errors.UnWrapResponse(errors.ErrInternalServer)
			res.ERR = err
		}
	} else {
		res = errors.UnWrapResponse(errors.ErrInternalServer)
	}
	if len(status) > 0 {
		res.Status = status[0]
	}
	if err := res.ERR; err != nil {
		if res.Message == "" {
			res.Message = err.Error()
		}
		if status := res.Status; status >= 400 && status < 500 {

		} else if status >= 500 {

		}
	}
	eitem := schema.ErrorItem{
		Code:    res.Code,
		Message: res.Message,
	}
	ResJSON(c, res.Status, schema.ErrorResult{Error: eitem})

}
