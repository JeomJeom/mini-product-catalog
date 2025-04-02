package service_test

import (
	"database/sql"
	"errors"
	"reflect"
	"testing"

	"thalesapi/data/dto"
	"thalesapi/data/models"
	"thalesapi/data/params"
	"thalesapi/db/builder"
	"thalesapi/internal/repository/mocks"
	"thalesapi/internal/service"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

// TestAddProduct verifies that the service calls CreateProduct on the repository.
func TestProductService_AddProduct(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockProductRepo(ctrl)
	prodService := service.NewProductService(mockRepo)

	// Dummy input attributes.
	productAttrs := &dto.ProductMutableAttrs{
		// populate fields as required for testing
		Name:         "Test Product",
		Price:        100.0,
		ModelNo:      "TP-123",
		Year:         2023,
		ThemeType:    "Test Theme",
		CategoryType: "Test Category",
		ImageURL:     "http://example.com/image.jpg",
		Description:  "Test description",
	}

	// Expect that CreateProduct is called with any *models.Product (we can use gomock.Any())
	mockRepo.EXPECT().CreateProduct(gomock.Any()).Return(nil)

	// Call service.
	err := prodService.AddProduct(productAttrs)
	assert.Nil(t, err, "AddProduct should not return an error")
}

// TestGetProduct tests both success and error scenarios for GetProduct.
func TestProductService_GetProduct(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockProductRepo(ctrl)
	prodService := service.NewProductService(mockRepo)

	// Success case.
	expectedProduct := &models.Product{
		ID:           "123",
		Name:         "Test Product",
		ModelNo:      "TP-123",
		Year:         2023,
		ThemeType:    "Test Theme",
		CategoryType: "Test Category",
		ImageURL:     "http://example.com/image.jpg",
		Price:        100.0,
		Description: sql.NullString{
			String: "Test description",
			Valid:  true,
		},
	}
	expectedView := expectedProduct.ToView()
	mockRepo.EXPECT().FindProduct("123").Return(expectedProduct, nil)

	view, err := prodService.GetProduct("123")
	assert.Nil(t, err, "GetProduct should not return an error")

	if !reflect.DeepEqual(view, expectedView) {
		assert.Fail(t, "GetProduct returned unexpected view", "Expected: %+v, Got: %+v", expectedView, view)
	}

	// Error case.
	mockRepo.EXPECT().FindProduct("not-found").Return(nil, errors.New("not found"))
	_, err = prodService.GetProduct("not-found")
	if err == nil {
		assert.Error(t, err, "Expected error for non-existent product, got nil")
	}
}

// TestGetProducts verifies that GetProducts returns the expected paginated result.
func TestProductService_GetProducts(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockProductRepo(ctrl)
	prodService := service.NewProductService(mockRepo)

	// Set up dummy filter and pageInfo.
	filter := &params.ProductFilter{
		Keyword:   "Test",
		ThemeType: "Test Theme",
	}
	pageInfo := &builder.PageInfo{
		PageIndex: 1,
		PageSize:  10,
	} // Example: page 1, 10 items per page.

	// Prepare a slice of dummy products.
	product1 := &models.Product{
		Name:         "Test Product 1",
		ModelNo:      "TP-1",
		Year:         2023,
		ThemeType:    "Test Theme",
		CategoryType: "Test Category",
		Price:        100.0,
		ImageURL:     "http://example.com/image1.jpg",
		Description: sql.NullString{
			String: "Test description 1",
			Valid:  true,
		},
	}
	product2 := &models.Product{
		Name:         "Test Product 2",
		ModelNo:      "TP-2",
		Year:         2023,
		ThemeType:    "Test Theme",
		CategoryType: "Test Category",
		Price:        200.0,
		ImageURL:     "http://example.com/image2.jpg",
		Description: sql.NullString{
			String: "Test description 2",
			Valid:  true,
		},
	}
	expectedProducts := []*models.Product{product1, product2}

	// Expect the repository calls.
	mockRepo.EXPECT().FindProducts(filter, pageInfo).Return(expectedProducts, nil)
	mockRepo.EXPECT().CountTotalProducts(filter, pageInfo).Return(int64(len(expectedProducts)), nil)

	result, err := prodService.GetProducts(filter, pageInfo)
	if err != nil {
		assert.Fail(t, "GetProducts returned an unexpected error", err)
	}

	if len(result.Data) != len(expectedProducts) {
		assert.Fail(t, "GetProducts returned unexpected number of products", "Expected: %d, Got: %d", len(expectedProducts), len(result.Data))
	}
}

// TestUpdateProduct tests that UpdateProduct calls the repository correctly.
func TestProductService_UpdateProduct(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockProductRepo(ctrl)
	prodService := service.NewProductService(mockRepo)

	// Dummy new values.
	newValues := &dto.ProductMutableAttrs{
		Name:         "Updated Product",
		ModelNo:      "UP-123",
		Year:         2024,
		ThemeType:    "Updated Theme",
		CategoryType: "Updated Category",
		ImageURL:     "http://example.com/updated_image.jpg",
		Price:        150.0,
		Description:  "Updated description",
	}
	// Expected updated product.
	updatedProduct := &models.Product{
		// populate with expected updated data
		ID:           "123",
		Name:         newValues.Name,
		ModelNo:      newValues.ModelNo,
		Year:         newValues.Year,
		ThemeType:    newValues.ThemeType,
		CategoryType: newValues.CategoryType,
		ImageURL:     newValues.ImageURL,
		Price:        newValues.Price,
		Description: sql.NullString{
			String: newValues.Description,
			Valid:  newValues.Description != "",
		},
	}

	mockRepo.EXPECT().UpdateProduct("123", newValues).Return(updatedProduct, nil)

	result, err := prodService.UpdateProduct("123", newValues)
	if err != nil {
		assert.Fail(t, "UpdateProduct returned an unexpected error", err)
	}
	if !reflect.DeepEqual(result, updatedProduct) {
		assert.Fail(t, "UpdateProduct returned unexpected product", "Expected: %+v, Got: %+v", updatedProduct, result)
	}
}

// TestDeleteProduct ensures DeleteProduct propagates errors from the repository.
func TestProductService_DeleteProduct(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockProductRepo(ctrl)
	prodService := service.NewProductService(mockRepo)

	// Success case.
	mockRepo.EXPECT().DeleteProductByID("123").Return(nil)
	err := prodService.DeleteProduct("123")
	if err != nil {
		assert.Fail(t, "DeleteProduct returned an unexpected error", err)
	}

	// Error case.
	mockRepo.EXPECT().DeleteProductByID("not-found").Return(errors.New("delete error"))
	err = prodService.DeleteProduct("not-found")
	if err == nil {
		assert.Nil(t, err, "Expected error when deleting non-existent product, got nil")
	}
}
