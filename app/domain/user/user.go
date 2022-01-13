package user

import (
	"demoddd/app/domain/pagination"
	"time"
)

type User struct {
	ID        string
	UserName  string
	RealName  string
	Password  string
	Email     *string
	Phone     *string
	Status    int
	Creator   string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
	// Roles role.Roles

}

type Users []*User

type QueryParams struct {
	PaginationParam pagination.Param
	OrderFields     pagination.OrderFields
	UserName        string
	QueryValue      string
	Status          int
	RoleIDs         []string
}

//获取用户id集合
func (a Users) ToIDs() []string {
	idList := make([]string, len(a))
	for i, item := range a {
		idList[i] = item.ID
	}
	return idList
}

// func (a Users) FIllRoles(userRoles map[string])
