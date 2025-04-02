package service_test

import (
	"errors"
	"reflect"
	"testing"

	"thalesapi/internal/repository/mocks"
	"thalesapi/internal/service"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestThemeTypeService_GetThemeTypes_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Create a mock of the ThemeTypeRepo.
	mockRepo := mocks.NewMockThemeTypeRepo(ctrl)
	themeSvc := service.NewThemeTypeService(mockRepo)

	// Define the expected result.
	expectedTypes := []string{"light", "dark"}

	// Set expectation on the mock: GetThemeTypes will be called and return expectedTypes with no error.
	mockRepo.EXPECT().GetThemeTypes().Return(expectedTypes, nil)

	// Call the service method.
	result, err := themeSvc.GetThemeTypes()
	if err != nil {
		assert.Fail(t, "expected no error, got %v", err)
	}
	if !reflect.DeepEqual(result, expectedTypes) {
		assert.Fail(t, "expected %v, got %v", expectedTypes, result)
	}
}

func TestThemeTypeService_GetThemeTypes_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockThemeTypeRepo(ctrl)
	themeSvc := service.NewThemeTypeService(mockRepo)

	// Define an error that the repository will return.
	expectedErr := errors.New("database error")

	// Set expectation: when GetThemeTypes is called, it returns nil and the expected error.
	mockRepo.EXPECT().GetThemeTypes().Return(nil, expectedErr)

	// Call the service method.
	result, err := themeSvc.GetThemeTypes()
	if err == nil {
		assert.Fail(t, "expected an error, got nil")
	}
	if err.Error() != expectedErr.Error() {
		assert.Fail(t, "expected error %v, got %v", expectedErr, err)
	}
	if result != nil {
		assert.Fail(t, "expected nil result, got %v", result)
	}
}
