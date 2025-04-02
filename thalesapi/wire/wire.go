//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/google/wire"
	"gorm.io/gorm"
	"thalesapi/internal/controller"
	"thalesapi/internal/repository"
	"thalesapi/internal/service"
)

func InitProductService(db *gorm.DB) service.ProductService {
	panic(wire.Build(repository.NewProductRepository, service.NewProductService))
}

func InitProductController(db *gorm.DB) *controller.ProductController {
	panic(wire.Build(InitProductService, controller.NewProductController))
}

func InitThemeTypeService(db *gorm.DB) service.ThemeTypeService {
	panic(wire.Build(repository.NewThemeTypeRepository, service.NewThemeTypeService))
}

func InitThemeTypeController(db *gorm.DB) *controller.ThemeTypeController {
	panic(wire.Build(InitThemeTypeService, controller.NewThemeTypeController))
}
