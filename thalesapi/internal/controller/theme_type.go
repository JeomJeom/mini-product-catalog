package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"thalesapi/internal/service"
)

type ThemeTypeController struct {
	svc service.ThemeTypeService
}

func NewThemeTypeController(svc service.ThemeTypeService) *ThemeTypeController {
	return &ThemeTypeController{svc: svc}
}

func (p *ThemeTypeController) GetThemeTypes(c *gin.Context) {
	themeTypes, err := p.svc.GetThemeTypes()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not fetch products"})
		return
	}
	c.JSON(http.StatusOK, themeTypes)
}
