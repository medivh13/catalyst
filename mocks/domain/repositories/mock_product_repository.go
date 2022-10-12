package mock_repository

/*
 * Author      : Jody (jody.almaida@gmail)
 * Modifier    :
 * Domain      : catalyst
 */

import (
	dto "catalyst/src/app/dtos/product"
	"catalyst/src/infra/models"

	"github.com/stretchr/testify/mock"
)

type MockProductRepo struct {
	mock.Mock
}

func (o *MockProductRepo) CreateProduct(dataProduct *dto.CreateProductReqDTO) error {
	args := o.Called(dataProduct)

	var (
		err error
	)

	if n, ok := args.Get(0).(error); ok {
		err = n
	}

	return err
}

func (o *MockProductRepo) GetSingleProduct(dataProduct *dto.GetProductReqDTO) (*models.Products, error) {
	args := o.Called(dataProduct)

	var (
		err  error
		data *models.Products
	)

	if n, ok := args.Get(0).(*models.Products); ok {
		data = n
	}

	if n, ok := args.Get(1).(error); ok {
		err = n
	}

	return data, err
}

func (o *MockProductRepo) GetProductByBrandId(dataProduct *dto.GetProductReqDTO) ([]*models.Products, error) {
	args := o.Called(dataProduct)

	var (
		err  error
		data []*models.Products
	)

	if n, ok := args.Get(0).([]*models.Products); ok {
		data = n
	}

	if n, ok := args.Get(1).(error); ok {
		err = n
	}

	return data, err
}
