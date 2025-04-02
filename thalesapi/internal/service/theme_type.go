package service

import (
	"thalesapi/internal/repository"
)

type (
	ThemeTypeService interface {
		GetThemeTypes() ([]string, error)
	}
	themeTypeSvcImpl struct {
		repo repository.ThemeTypeRepo
	}
)

func NewThemeTypeService(repo repository.ThemeTypeRepo) ThemeTypeService {
	return &themeTypeSvcImpl{repo: repo}
}

func (t themeTypeSvcImpl) GetThemeTypes() ([]string, error) {
	return t.repo.GetThemeTypes()
}
