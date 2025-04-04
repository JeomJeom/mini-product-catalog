// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package wire

import (
	"gorm.io/gorm"
	"thalesapi/internal/controller"
	"thalesapi/internal/repository"
	"thalesapi/internal/service"
)

// Injectors from wire.go:

func InitProductService(db *gorm.DB) service.ProductService {
	productRepo := repository.NewProductRepository(db)
	productService := service.NewProductService(productRepo)
	return productService
}

func InitProductController(db *gorm.DB) *controller.ProductController {
	productService := InitProductService(db)
	productController := controller.NewProductController(productService)
	return productController
}

func InitThemeTypeService(db *gorm.DB) service.ThemeTypeService {
	themeTypeRepo := repository.NewThemeTypeRepository(db)
	themeTypeService := service.NewThemeTypeService(themeTypeRepo)
	return themeTypeService
}

func InitThemeTypeController(db *gorm.DB) *controller.ThemeTypeController {
	themeTypeService := InitThemeTypeService(db)
	themeTypeController := controller.NewThemeTypeController(themeTypeService)
	return themeTypeController
}
