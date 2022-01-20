package schema

import (
	"thrgoweb/pkg/util/json"
	"time"
)

//处理表单的工具

type User struct {
	ID        uint64    `json:"id,string"`
	UserName  string    `json:"user_name" binding:"required"`
	RealName  string    `json:"real_name" binding:"required"`
	Password  string    `json:"password"`
	Phone     string    `json:"phone"`
	Email     string    `json:"email"`
	Status    int       `json:"status" binding:"required,max=2,min=1"`
	Creator   uint64    `json:"creator"`
	CreatedAt time.Time `json:"created_at"`
	UserRoles UserRoles `json:"user_roles" binding:"required,gt=0"`
}

type UserRole struct {
	ID     uint64 `json:"id,string"`
	UserID uint64 `json:"user_id,string"`
	RoleID uint64 `json:"role_id,string"`
}

// UserRoles 角色菜单列表
type UserRoles []*UserRole

func (a *User) String() string {
	return json.MarshalToString(a)
}

// CleanSecure 清理安全数据
func (a *User) CleanSecure() *User {
	a.Password = ""
	return a
}

// UserQueryParam 查询条件
type UserQueryParam struct {
	PaginationParam
	UserName   string   `form:"userName"`
	QueryValue string   `form:"queryValue"`
	Status     int      `form:"status"`
	RoleIDs    []uint64 `form:"-"`
}

// UserQueryOptions 查询可选参数项
type UserQueryOptions struct {
	OrderFields  []*OrderField
	SelectFields []string
}

type Users []*User

// UserQueryResult 查询结果
type UserQueryResult struct {
	Data       Users
	PageResult *PaginationResult
}

// ToShowResult 转换为显示结果

// ToUserShows 转换为用户显示列表

// ToIDs 转换为唯一标识列表
func (a Users) ToIDs() []uint64 {
	idList := make([]uint64, len(a))
	for i, item := range a {
		idList[i] = item.ID
	}
	return idList
}
