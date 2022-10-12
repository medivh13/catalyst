package mock_dto

import (
	dto "catalyst/src/app/dtos/brand"

	"github.com/stretchr/testify/mock"
)

type MockBrandDTO struct {
	mock.Mock
}

func NewMockBrandDTO() *MockBrandDTO {
	return &MockBrandDTO{}
}

var _ dto.BrandInterface = &MockBrandDTO{}

func (m *MockBrandDTO) Validate() error {
	args := m.Called()
	var err error
	if n, ok := args.Get(0).(error); ok {
		err = n
		return err
	}

	return nil
}
