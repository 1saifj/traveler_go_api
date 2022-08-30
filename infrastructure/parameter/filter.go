package parameter

type Filter struct {
	Limit      *int    `json:"limit" form:"limit"`
	Offset     *int    `json:"offset" form:"offset"`
	OrderBy    *string `url:"order_by"`
	OrderField *string `url:"order_field"`
}

func (filter *Filter) GetLimit() int {
	if filter.Limit == nil {
		return 20
	}
	return *filter.Limit
}

func (filter *Filter) GetOffset() int {
	if filter.Offset == nil {
		return 0
	}
	return *filter.Offset
}

func (filter *Filter) OrderQueryBy() string {
	if filter.OrderBy == nil || filter.OrderField == nil {
		return `created_at desc`
	} else {
		return *filter.OrderBy + *filter.OrderField
	}
}
