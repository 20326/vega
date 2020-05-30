package model


type (
	PageQuery struct {
		Where     string
		WhereArgs []interface{}
		PageNo    int
		PageSize  int
	}
)
