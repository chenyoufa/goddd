package schema

import "time"

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
	// UserRoles UserRoles `json:"id,string"`
}

// UserQueryParam 查询条件
type UserQueryParam struct {
	// PaginationParam
	UserName   string   `form:"userName"`   // 用户名
	QueryValue string   `form:"queryValue"` // 模糊查询
	Status     int      `form:"status"`     // 用户状态(1:启用 2:停用)
	RoleIDs    []uint64 `form:"-"`          // 角色ID列表
}

// UserQueryOptions 查询可选参数项
type UserQueryOptions struct {
	// OrderFields  []*OrderField
	SelectFields []string
}
