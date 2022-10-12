package mock_dto

import (
	dto "catalyst/src/app/dtos/order_detail"

	"github.com/stretchr/testify/mock"
)

type MockOrderDetailDTO struct {
	mock.Mock
}

func NewMockOrderDTO() *MockOrderDetailDTO {
	return &MockOrderDetailDTO{}
}

var _ dto.OrderDetailInterface = &MockOrderDetailDTO{}

func (m *MockOrderDetailDTO) Validate() error {
	args := m.Called()
	var err error
	if n, ok := args.Get(0).(error); ok {
		err = n
		return err
	}

	return nil
}
