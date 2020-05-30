package model


type (
	PageQuery struct {
		Where     string
		WhereArgs []interface{}
		PageNo    int
		PageSize  int
	}
)

func NewPageQuery(pageNo int, pageSize int)  PageQuery {
	return PageQuery{
		Where: "",
		WhereArgs: []interface{}{},
		PageNo: pageNo,
		PageSize: pageSize,
	}
}