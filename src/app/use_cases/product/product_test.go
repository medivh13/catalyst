package product_usecases

/*
 * Author      : Jody (jody.almaida@gmail.com)
 * Modifier    :
 * Domain      : catalyst
 */

import (
	mockDTO "catalyst/mocks/app/dtos/product"
	mockRepo "catalyst/mocks/domain/repositories"
	"catalyst/src/infra/models"
	"errors"
	"testing"

	dto "catalyst/src/app/dtos/product"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type MockUsecase struct {
	mock.Mock
}

type UsecaseProducttTest struct {
	suite.Suite
	repo                 *mockRepo.MockProductRepo
	models               *models.Products
	listModels           []*models.Products
	usecase              ProductUCInterface
	dtoTest              *dto.CreateProductReqDTO
	dtoTestSingleProduct *dto.GetProductReqDTO
	mockDTO              *mockDTO.MockProductDTO
}

func (suite *UsecaseProducttTest) SetupTest() {
	suite.repo = new(mockRepo.MockProductRepo)
	suite.mockDTO = new(mockDTO.MockProductDTO)
	suite.usecase = NewProductUseCase(suite.repo)

	suite.dtoTest = &dto.CreateProductReqDTO{
		Name: "Test1",
	}

	suite.dtoTestSingleProduct = &dto.GetProductReqDTO{
		ID: 1,
	}

	suite.models = &models.Products{
		ID:        1,
		Name:      "test1",
		Price:     1000,
		BrandID:   1,
		BrandName: "test1",
	}

	suite.listModels = []*models.Products{
		&models.Products{
			ID:        1,
			Name:      "test1",
			Price:     1000,
			BrandID:   1,
			BrandName: "test1",
		},
	}
}

func (uc *UsecaseProducttTest) TestCreateProductSuccess() {
	uc.repo.Mock.On("CreateProduct", uc.dtoTest).Return(nil)
	err := uc.usecase.CreateProduct(uc.dtoTest)
	uc.Equal(nil, err)
}

func (uc *UsecaseProducttTest) TestCreateProductFail() {
	uc.repo.Mock.On("CreateProduct", uc.dtoTest).Return(errors.New(mock.Anything))
	err := uc.usecase.CreateProduct(uc.dtoTest)
	uc.Equal(errors.New(mock.Anything), err)
}

func (uc *UsecaseProducttTest) TestGetSingleProductSuccess() {
	uc.repo.Mock.On("GetSingleProduct", uc.dtoTestSingleProduct).Return(uc.models, nil)
	_, err := uc.usecase.GetSingleProduct(uc.dtoTestSingleProduct)
	uc.Equal(nil, err)
}

func (uc *UsecaseProducttTest) TestGetSingleProductFail() {
	uc.repo.Mock.On("GetSingleProduct", uc.dtoTestSingleProduct).Return(nil, errors.New(mock.Anything))
	_, err := uc.usecase.GetSingleProduct(uc.dtoTestSingleProduct)
	uc.Equal(errors.New(mock.Anything), err)
}

func (uc *UsecaseProducttTest) TestGetProductByBrandIDSuccess() {
	uc.repo.Mock.On("GetProductByBrandId", uc.dtoTestSingleProduct).Return(uc.listModels, nil)
	_, err := uc.usecase.GetProductByBrandID(uc.dtoTestSingleProduct)
	uc.Equal(nil, err)
}

func (uc *UsecaseProducttTest) TestGetProductByBrandIDFail() {
	uc.repo.Mock.On("GetProductByBrandId", uc.dtoTestSingleProduct).Return(nil, errors.New(mock.Anything))
	_, err := uc.usecase.GetProductByBrandID(uc.dtoTestSingleProduct)
	uc.Equal(errors.New(mock.Anything), err)
}

func TestUsecase(t *testing.T) {
	suite.Run(t, new(UsecaseProducttTest))
}
