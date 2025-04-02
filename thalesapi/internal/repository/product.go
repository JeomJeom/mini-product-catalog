package repository

import (
	"thalesapi/data/dto"
	"thalesapi/data/models"
	"thalesapi/data/params"
	"thalesapi/db/builder"

	"github.com/doug-martin/goqu/v9"
	"github.com/doug-martin/goqu/v9/exp"
	"gorm.io/gorm"
)

type (
	ProductRepo interface {
		CreateProduct(product *models.Product) error
		FindProducts(filter *params.ProductFilter, pageInfo *builder.PageInfo) ([]*models.Product, error)
		FindProduct(id string) (*models.Product, error)
		UpdateProduct(id string, newValues *dto.ProductMutableAttrs) (*models.Product, error)
		DeleteProductByID(id string) error
		CountTotalProducts(filter *params.ProductFilter, pageInfo *builder.PageInfo) (int64, error)
	}
	ProductRepoImpl struct {
		DB      *gorm.DB
		Builder *goqu.DialectWrapper
	}
)

func NewProductRepository(db *gorm.DB) ProductRepo {

	return &ProductRepoImpl{DB: db, Builder: builder.NewDialectWrapper(db)}
}

func (p ProductRepoImpl) CreateProduct(product *models.Product) error {
	return p.DB.Create(product).Error
}

func (p ProductRepoImpl) FindProducts(filter *params.ProductFilter, pageInfo *builder.PageInfo) ([]*models.Product, error) {
	var products []*models.Product

	stmt := p.Builder.From("products").
		Where(filter.ApplyWhereExp([]exp.Expression{goqu.C("deleted_at").IsNull()})...)
	stmt = pageInfo.BuildPaginatedQuery(stmt)
	sqlStmt, _, _ := stmt.ToSQL()

	return products, p.DB.Raw(sqlStmt).Scan(&products).Error
}

func (p ProductRepoImpl) FindProduct(id string) (*models.Product, error) {
	var product models.Product
	err := p.DB.Where("id = ?", id).First(&product).Error
	if err != nil {
		return nil, err
	}
	return &product, err
}

func (p ProductRepoImpl) UpdateProduct(id string, newValues *dto.ProductMutableAttrs) (*models.Product, error) {
	var product models.Product
	product.Modify(newValues)

	err := p.DB.Model(&product).Where("id = ?", id).Updates(&product).Error
	if err != nil {
		return nil, err
	}
	return &product, err
}

func (p ProductRepoImpl) CountTotalProducts(filter *params.ProductFilter, pageInfo *builder.PageInfo) (int64, error) {
	var count int64
	stmt := p.Builder.From("products").
		Where(filter.ApplyWhereExp([]exp.Expression{})...).
		Select(goqu.COUNT("id"))

	sqlStmt, _, _ := stmt.ToSQL()
	return count, p.DB.Raw(sqlStmt).Scan(&count).Error
}

func (p ProductRepoImpl) DeleteProductByID(id string) error {
	return p.DB.Delete(&models.Product{}, id).Error
}
