package order_usecases

import (
	mockDTO "catalyst/mocks/app/dtos/order"
	mockDTODetail "catalyst/mocks/app/dtos/order_detail"
	mockRepo "catalyst/mocks/domain/repositories"
	"errors"

	"catalyst/src/infra/models"
	"testing"

	dto "catalyst/src/app/dtos/order"
	dtoDetail "catalyst/src/app/dtos/order_detail"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

/*
 * Author      : Jody (jody.almaida@gmail.com)
 * Modifier    :
 * Domain      : catalyst
 */

type MockUsecase struct {
	mock.Mock
}

type UsecaseOrdertTest struct {
	suite.Suite
	repo          *mockRepo.MockOrderRepo
	models        *models.OrderReqModel
	listModels    []*models.Orders
	usecase       OrderUCInterface
	dtoTest       *dto.CreateOrderReqDTO
	dtoDetail     *dtoDetail.GetOrderReqDTO
	mockDTO       *mockDTO.MockOrderDTO
	mockDTODetail *mockDTODetail.MockOrderDetailDTO
}

func (suite *UsecaseOrdertTest) SetupTest() {
	suite.repo = new(mockRepo.MockOrderRepo)
	suite.mockDTO = new(mockDTO.MockOrderDTO)
	suite.usecase = NewOrderUseCase(suite.repo)

	suite.dtoTest = &dto.CreateOrderReqDTO{
		Data: []*dto.DetailOrderReqDTO{
			&dto.DetailOrderReqDTO{
				ProductID: 1,
				Price:     1000,
				Quantity:  1,
			},
		},
	}

	suite.dtoDetail = &dtoDetail.GetOrderReqDTO{
		OrderID: "1",
	}

	suite.models = &models.OrderReqModel{
		OrderID: 1,
	}

	suite.listModels = []*models.Orders{
		&models.Orders{
			ID:         1,
			OrderCode:  "test",
			Total:      2000,
			Product:    "test",
			Quantity:   2,
			TotalPrice: 1000,
		},
	}
}

func (uc *UsecaseOrdertTest) TestCreateOrderSuccess() {
	uc.repo.Mock.On("CreateOrder", uc.dtoTest).Return(uc.models, nil)
	_, err := uc.usecase.CreateOrder(uc.dtoTest)
	uc.Equal(nil, err)
}

func (uc *UsecaseOrdertTest) TestCreateOrderFail() {
	uc.repo.Mock.On("CreateOrder", uc.dtoTest).Return(nil, errors.New(mock.Anything))
	_, err := uc.usecase.CreateOrder(uc.dtoTest)
	uc.Equal(errors.New(mock.Anything), err)
}

func (uc *UsecaseOrdertTest) TestGetOrderByIDSuccess() {
	uc.repo.Mock.On("GetOrderByID", uc.dtoDetail).Return(uc.listModels, nil)
	_, err := uc.usecase.GetOrderByID(uc.dtoDetail)
	uc.Equal(nil, err)
}

func (uc *UsecaseOrdertTest) TestGetOrderByIDFail() {
	uc.repo.Mock.On("GetOrderByID", uc.dtoDetail).Return(nil, errors.New(mock.Anything))
	_, err := uc.usecase.GetOrderByID(uc.dtoDetail)
	uc.Equal(errors.New(mock.Anything), err)
}

func TestUsecase(t *testing.T) {
	suite.Run(t, new(UsecaseOrdertTest))
}
