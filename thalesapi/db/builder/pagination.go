package builder

import (
	"github.com/doug-martin/goqu/v9"
	"github.com/gin-gonic/gin"
	"strconv"
	"thalesapi/data/enums"
)

// PageInfo holds pagination and sorting parameters.
type PageInfo struct {
	PageIndex  uint          `json:"pageIndex"`
	PageSize   uint          `json:"pageSize"`
	ColumnName string        `json:"columnName"`
	OrderBy    enums.OrderBy `json:"orderBy"`
}

func NewPaginatedRequest(c *gin.Context, defaultColumn string) (*PageInfo, error) {
	page, _ := strconv.Atoi(c.DefaultQuery("pageIndex", "0"))
	size, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))
	column := c.DefaultQuery("columnName", defaultColumn)
	orderBy := c.DefaultQuery("orderBy", "desc")

	return &PageInfo{
		PageIndex:  uint(page),
		PageSize:   uint(size),
		ColumnName: column,
		OrderBy:    enums.ParseOrderBy(orderBy),
	}, nil
}

func ComputeOffset(page uint, size uint) uint {
	return page * size
}

// BuildPaginatedQuery applies offset, limit, and ordering to a goqu.SelectDataset based on PageInfo.
// If PageInfo is nil or ColumnName is empty, it returns the dataset without applying ordering.
func (p *PageInfo) BuildPaginatedQuery(stmt *goqu.SelectDataset) *goqu.SelectDataset {
	if p == nil {
		return stmt
	}

	// Calculate offset and limit.
	offset := ComputeOffset(p.PageIndex, p.PageSize)
	stmt = stmt.Offset(offset).Limit(p.PageSize)

	// If no column is provided, skip ordering.
	if p.ColumnName == "" {
		return stmt
	}

	// Apply ordering based on OrderBy value.
	column := goqu.C(p.ColumnName)
	if p.OrderBy == enums.Desc {
		switch p.ColumnName {
		case "updated_at":
			stmt = stmt.Order(column.Desc().NullsLast())
		default:
			stmt = stmt.Order(column.Desc())
		}
	} else {
		switch p.ColumnName {
		case "updated_at":
			stmt = stmt.Order(column.Asc().NullsLast())
		default:
			stmt = stmt.Order(column.Asc())
		}
	}

	return stmt
}
