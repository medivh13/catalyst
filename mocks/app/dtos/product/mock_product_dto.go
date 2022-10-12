package mock_dto

import (
	dto "catalyst/src/app/dtos/product"

	"github.com/stretchr/testify/mock"
)

type MockProductDTO struct {
	mock.Mock
}

func NewMockProductDTO() *MockProductDTO {
	return &MockProductDTO{}
}

var _ dto.PrroductInterface = &MockProductDTO{}

func (m *MockProductDTO) Validate() error {
	args := m.Called()
	var err error
	if n, ok := args.Get(0).(error); ok {
		err = n
		return err
	}

	return nil
}
