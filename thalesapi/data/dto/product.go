package dto

import (
	"time"
)

type (
	ProductView struct {
		ID           string    `json:"id"`
		CreatedAt    time.Time `json:"createdAt"`
		UpdatedAt    time.Time `json:"updatedAt"`
		Name         string    `json:"name"`
		ModelNo      string    `json:"modelNo"`
		Year         int       `json:"year"`
		ThemeType    string    `json:"themeType"`
		CategoryType string    `json:"categoryType"`
		ImageURL     string    `json:"imageURL"`
		Price        float64   `json:"price"`
		Description  string    `json:"description"`
	}

	PaginatedProductResp struct {
		Data       []*ProductView `json:"data"`
		Pagination Pagination     ` json:"pagination"`
	}

	ProductMutableAttrs struct {
		Name         string  `json:"name"`
		ModelNo      string  `json:"modelNo"`
		Year         int     `json:"year"`
		ThemeType    string  `json:"themeType"`
		CategoryType string  `json:"categoryType"`
		ImageURL     string  `json:"imageURL"`
		Price        float64 `json:"price"`
		Description  string  `json:"description"`
	}
)
