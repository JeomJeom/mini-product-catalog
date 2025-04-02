package dto

import (
	"math"
	"thalesapi/db/builder"
)

type Pagination struct {
	PageIndex  uint  `json:"pageIndex"`
	PageSize   uint  `json:"pageSize"`
	TotalItems int64 `json:"totalItems"`
	TotalPages uint  `json:"totalPages"`
}

func (p *Pagination) Assign(info *builder.PageInfo, totalCount int64) {
	p.PageSize = info.PageSize
	p.PageIndex = info.PageIndex
	p.TotalItems = totalCount
	p.TotalPages = uint(math.Ceil(float64(totalCount) / float64(info.PageSize)))
}
