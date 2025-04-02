package interfaces

import (
	"gorm.io/gorm"
)

// DB is a minimal interface for the GORM methods we use.
type DB interface {
	Create(v interface{}) *gorm.DB
	Where(query interface{}, args ...interface{}) *gorm.DB
	First(dest interface{}) *gorm.DB
	Debug() *gorm.DB
	Raw(sql string, values ...interface{}) *gorm.DB
	Delete(value interface{}, conds ...interface{}) *gorm.DB
	Model(value interface{}) *gorm.DB
	Updates(values interface{}) *gorm.DB
}

//func NewProductRepositoryWithDB(db DB) repository.ProductRepo {
//	// Builder can be built using the actual db (or a wrapped version of it).
//	return &repository.ProductRepoImpl{DB: db.(*gorm.DB), Builder: builder.NewDialectWrapper(db.(*gorm.DB))}
//}
