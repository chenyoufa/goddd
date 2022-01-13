package pagination

type Pagination struct {
	Total    int  `json:"total"`
	Current  uint `json:"current"`
	PageSize uint `json:"pageSize"`
}

type Param struct {
	Pagination bool `form:"-"`
	OnlyCount  bool `form:"-"`
	Current    uint `form:"current,default=1"`
	PageSize   uint `form:"pageSize,default=10" binding:"max=100"`
}

func (a Param) GetCurrent() uint {
	return a.Current
}
func (a Param) GetPageSize() uint {
	pageSize := a.PageSize
	if a.PageSize == 0 {
		pageSize = 100
	}
	return pageSize
}

type OrderDirection int

const (
	OrderByASC  OrderDirection = 1
	OrderByDESC OrderDirection = 2
)

type OrderField struct {
	Key       string
	Direction OrderDirection
}

func NewOrderFields(orderFields ...*OrderField) []*OrderField {
	return orderFields
}

func NewOrderField(key string, d OrderDirection) *OrderField {
	return &OrderField{
		Key:       key,
		Direction: d,
	}
}

type OrderFields []*OrderField

func (a OrderFields) AddIdSortField() OrderFields {
	return append(a, NewOrderField("id", OrderByDESC))
}
