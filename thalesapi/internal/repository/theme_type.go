package repository

import (
	"gorm.io/gorm"
	"thalesapi/data/models"
)

type (
	ThemeTypeRepo interface {
		GetThemeTypes() ([]string, error)
	}
	themeTypeRepoImpl struct {
		db *gorm.DB
	}
)

func NewThemeTypeRepository(db *gorm.DB) ThemeTypeRepo {
	return &themeTypeRepoImpl{db: db}
}

func (t themeTypeRepoImpl) GetThemeTypes() ([]string, error) {
	var values []string

	return values, t.db.Model([]models.Product{}).Distinct("theme_type").Pluck("theme_type", &values).Error
}
