package mock_dto

import (
	dto "catalyst/src/app/dtos/order"

	"github.com/stretchr/testify/mock"
)

type MockOrderDTO struct {
	mock.Mock
}

func NewMockOrderDTO() *MockOrderDTO {
	return &MockOrderDTO{}
}

var _ dto.OrderInterface = &MockOrderDTO{}

func (m *MockOrderDTO) Validate() error {
	args := m.Called()
	var err error
	if n, ok := args.Get(0).(error); ok {
		err = n
		return err
	}

	return nil
}
