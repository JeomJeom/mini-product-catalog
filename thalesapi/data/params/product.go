package params

import (
	"github.com/doug-martin/goqu/v9"
	"github.com/doug-martin/goqu/v9/exp"
	"github.com/gin-gonic/gin"
	"strings"
)

type ProductFilter struct {
	Keyword   string `json:"keyword"`
	ThemeType string `json:"themeType"`
}

func NewProductFilter(c *gin.Context) *ProductFilter {
	return &ProductFilter{Keyword: c.Query("keyword"), ThemeType: c.Query("themeType")}
}

func (p *ProductFilter) ApplyWhereExp(whereEx []exp.Expression) []exp.Expression {
	if p != nil {
		if p.Keyword != "" {
			whereEx = append(whereEx, goqu.L("LOWER(name)").Like("%"+strings.ToLower(p.Keyword)+"%"))
		}

		if p.ThemeType != "" {
			whereEx = append(whereEx, goqu.C("theme_type").Eq(p.ThemeType))
		}
	}

	return whereEx
}
