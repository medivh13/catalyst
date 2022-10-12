package brand_usecases

/*
 * Author      : Jody (jody.almaida@gmail.com)
 * Modifier    :
 * Domain      : catalyst
 */

import (
	mockDTO "catalyst/mocks/app/dtos/brand"
	mockRepo "catalyst/mocks/domain/repositories"
	"errors"
	"testing"

	dto "catalyst/src/app/dtos/brand"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type MockUsecase struct {
	mock.Mock
}

type UsecaseBrandtTest struct {
	suite.Suite
	repo *mockRepo.MockBrandRepo

	usecase BrandUCInterface
	dtoTest *dto.CreateBrandReqDTO
	mockDTO *mockDTO.MockBrandDTO
}

func (suite *UsecaseBrandtTest) SetupTest() {
	suite.repo = new(mockRepo.MockBrandRepo)
	suite.mockDTO = new(mockDTO.MockBrandDTO)
	suite.usecase = NewBrandUseCase(suite.repo)

	suite.dtoTest = &dto.CreateBrandReqDTO{
		Name: "Test1",
	}

}

func (uc *UsecaseBrandtTest) TestCreateBrandSuccess() {
	uc.repo.Mock.On("CreateBrand", uc.dtoTest).Return(nil)
	err := uc.usecase.CreateBrand(uc.dtoTest)
	uc.Equal(nil, err)
}

func (uc *UsecaseBrandtTest) TestCreateBrandFail() {
	uc.repo.Mock.On("CreateBrand", uc.dtoTest).Return(errors.New(mock.Anything))
	err := uc.usecase.CreateBrand(uc.dtoTest)
	uc.Equal(errors.New(mock.Anything), err)
}

func TestUsecase(t *testing.T) {
	suite.Run(t, new(UsecaseBrandtTest))
}
