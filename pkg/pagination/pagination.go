package pagination

import (
	"math"
)

// Pagination represents pagination info.
type Pagination map[string]interface{}

func (p Pagination) NewPagination(pageNo, pageSize, totalCount int) Pagination {
	pageCount := int(math.Ceil(float64(totalCount) / float64(pageSize)))
	return Pagination{
		"pageNo":     pageNo,
		"pageSize":   pageSize,
		"pageCount":  pageCount,
		"totalCount": totalCount,
	}
}

func (p Pagination) SetData(data interface{}) Pagination {
	p["data"] = data
	return p
}

//// NewPagination creates a new pagination with the specified current page num, page size, window size and record count.
//func NewPagination(pageNo, pageSize, totalCount int) Pagination {
//	pageCount := int(math.Ceil(float64(totalCount) / float64(pageSize)))
//
//	return Pagination{
//		"pageNo":     pageNo,
//		"pageSize":   pageSize,
//		"pageCount":  pageCount,
//		"totalCount": totalCount,
//	}
//}
