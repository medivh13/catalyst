package mock_repository

/*
 * Author      : Jody (jody.almaida@gmail)
 * Modifier    :
 * Domain      : catalyst
 */

import (
	dto "catalyst/src/app/dtos/brand"

	"github.com/stretchr/testify/mock"
)

type MockBrandRepo struct {
	mock.Mock
}

func (o *MockBrandRepo) CreateBrand(dataBrand *dto.CreateBrandReqDTO) error {
	args := o.Called(dataBrand)

	var (
		err error
	)

	if n, ok := args.Get(0).(error); ok {
		err = n
	}

	return err
}
