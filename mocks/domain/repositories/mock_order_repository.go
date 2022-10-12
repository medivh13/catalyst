package mock_repository

/*
 * Author      : Jody (jody.almaida@gmail)
 * Modifier    :
 * Domain      : catalyst
 */

import (
	dto "catalyst/src/app/dtos/order"
	dtoDetail "catalyst/src/app/dtos/order_detail"
	"catalyst/src/infra/models"

	"github.com/stretchr/testify/mock"
)

type MockOrderRepo struct {
	mock.Mock
}

func (o *MockOrderRepo) CreateOrder(dataOrder *dto.CreateOrderReqDTO) (*models.OrderReqModel, error) {
	args := o.Called(dataOrder)

	var (
		err  error
		data *models.OrderReqModel
	)

	if n, ok := args.Get(0).(*models.OrderReqModel); ok {
		data = n
	}

	if n, ok := args.Get(1).(error); ok {
		err = n
	}

	return data, err
}

func (o *MockOrderRepo) GetOrderByID(dataOrder *dtoDetail.GetOrderReqDTO) ([]*models.Orders, error) {
	args := o.Called(dataOrder)

	var (
		err  error
		data []*models.Orders
	)

	if n, ok := args.Get(0).([]*models.Orders); ok {
		data = n
	}

	if n, ok := args.Get(1).(error); ok {
		err = n
	}

	return data, err
}
