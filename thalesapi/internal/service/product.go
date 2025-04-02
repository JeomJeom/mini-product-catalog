package service

import (
	"thalesapi/data/dto"
	"thalesapi/data/models"
	"thalesapi/data/params"
	"thalesapi/db/builder"
	"thalesapi/internal/repository"
)

type (
	ProductService interface {
		AddProduct(product *dto.ProductMutableAttrs) error
		GetProduct(id string) (*dto.ProductView, error)
		GetProducts(filter *params.ProductFilter, pageInfo *builder.PageInfo) (*dto.PaginatedProductResp, error)
		UpdateProduct(id string, newValues *dto.ProductMutableAttrs) (*models.Product, error)
		DeleteProduct(id string) error
	}
	productServiceImpl struct {
		repo repository.ProductRepo
	}
)

func NewProductService(repo repository.ProductRepo) ProductService {
	return &productServiceImpl{repo: repo}
}

func (p productServiceImpl) AddProduct(product *dto.ProductMutableAttrs) error {
	var newProduct models.Product
	newProduct.Modify(product)
	return p.repo.CreateProduct(&newProduct)
}

func (p productServiceImpl) GetProduct(id string) (*dto.ProductView, error) {
	product, err := p.repo.FindProduct(id)
	if err != nil {
		return nil, err
	}

	return product.ToView(), nil
}

func (p productServiceImpl) GetProducts(filter *params.ProductFilter, pageInfo *builder.PageInfo) (*dto.PaginatedProductResp, error) {
	products, err := p.repo.FindProducts(filter, pageInfo)
	if err != nil {
		return nil, err
	}

	views := make([]*dto.ProductView, 0, len(products))
	for _, item := range products {
		views = append(views, item.ToView())
	}

	count, err := p.repo.CountTotalProducts(filter, pageInfo)
	if err != nil {
		return nil, err
	}

	result := &dto.PaginatedProductResp{
		Data: views,
	}
	result.Pagination.Assign(pageInfo, count)
	return result, nil
}

func (p productServiceImpl) UpdateProduct(id string, newValues *dto.ProductMutableAttrs) (*models.Product, error) {
	product, err := p.repo.UpdateProduct(id, newValues)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (p productServiceImpl) DeleteProduct(id string) error {
	return p.repo.DeleteProductByID(id)
}
