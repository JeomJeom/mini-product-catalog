package router

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"thalesapi/configs"
	"thalesapi/db"
	"thalesapi/wire"
)

type Router struct {
	engine *gin.Engine
	db     *gorm.DB
}

func NewRouter(engine *gin.Engine, config *configs.ProjectConfig) *Router {
	d := db.InitDB(config.DBConfig)
	return &Router{engine: engine, db: d}
}

func (r *Router) LoadRouter() {
	router := r.engine

	// for health checking purpose
	router.GET("/ping", func(context *gin.Context) {
		context.JSON(http.StatusOK, "OK")
	})

	mainApi := router.Group("/api/v1")
	mainApi.Use(func(context *gin.Context) {})
	{
		productApi := mainApi.Group("/products")
		productApi.Use(func(context *gin.Context) {})
		{
			productController := wire.InitProductController(r.db)
			productApi.POST("", productController.CreateProduct)
			productApi.GET("", productController.GetAllProducts)
			productApi.GET("/:id", productController.GetProduct)
			productApi.PUT("/:id", productController.UpdateProduct)
			productApi.DELETE("/:id", productController.DeleteProduct)
		}

		themeTypeApi := mainApi.Group("/theme-types")
		themeTypeApi.Use(func(context *gin.Context) {})
		{
			themeTypeController := wire.InitThemeTypeController(r.db)
			themeTypeApi.GET("", themeTypeController.GetThemeTypes)
		}
	}
}
