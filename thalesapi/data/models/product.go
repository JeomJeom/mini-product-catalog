package models

import (
	"database/sql"
	"thalesapi/data/dto"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Product struct {
	ID           string `gorm:"primarykey"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    sql.NullTime   `default:"NULL"`
	Name         string         `json:"name"`
	ModelNo      string         `json:"modelNo"`
	Year         int            `json:"year"`
	ThemeType    string         `json:"themeType"`
	CategoryType string         `json:"categoryType"`
	ImageURL     string         `json:"imageURL"`
	Price        float64        `json:"price"`
	Description  sql.NullString `json:"description"`
}

// BeforeCreate will set a UUID rather than numeric ID.
func (p *Product) BeforeCreate(tx *gorm.DB) (err error) {
	p.ID = uuid.New().String()
	p.CreatedAt = time.Now().UTC()
	p.UpdatedAt = time.Now().UTC()
	p.DeletedAt = sql.NullTime{
		Time:  time.Time{},
		Valid: false,
	}
	return
}

func (p *Product) Modify(view *dto.ProductMutableAttrs) {
	p.Name = view.Name
	p.CategoryType = view.CategoryType
	p.ImageURL = view.ImageURL
	p.Price = view.Price
	p.ModelNo = view.ModelNo
	p.Year = view.Year
	p.ThemeType = view.ThemeType
	p.UpdatedAt = time.Now().UTC()

	p.Description = sql.NullString{
		String: view.Description,
		Valid:  view.Description != "",
	}

}

func (p *Product) ToView() *dto.ProductView {
	return &dto.ProductView{
		ID:           p.ID,
		CreatedAt:    p.CreatedAt,
		UpdatedAt:    p.UpdatedAt,
		Name:         p.Name,
		ModelNo:      p.ModelNo,
		Year:         p.Year,
		ThemeType:    p.ThemeType,
		CategoryType: p.CategoryType,
		ImageURL:     p.ImageURL,
		Price:        p.Price,
		Description:  p.Description.String,
	}
}
