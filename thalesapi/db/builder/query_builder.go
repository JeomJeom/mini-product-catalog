package builder

import (
	"github.com/doug-martin/goqu/v9"
	"gorm.io/gorm"
)

func NewDialectWrapper(tx *gorm.DB) *goqu.DialectWrapper {
	dialect := goqu.Dialect(tx.Dialector.Name())
	return &dialect
}
