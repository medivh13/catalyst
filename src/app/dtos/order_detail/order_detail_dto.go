package order_detail_dto

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

type OrderDetailInterface interface {
	Validate() error
}

type GetOrderReqDTO struct {
	OrderID string `json:"id"`
}

func (dto *GetOrderReqDTO) Validate() error {
	if err := validation.ValidateStruct(
		dto,
		validation.Field(&dto.OrderID, validation.Required),
	); err != nil {
		return err
	}
	return nil
}

type GetOrderRespDTO struct {
	ID        int64                    `json:"id"`
	OrderCode string                   `json:"order_code"`
	Total     float64                  `json:"total"`
	Detail    []*GetOrderDetailRespDTO `json:"detail"`
}

type GetOrderDetailRespDTO struct {
	Product    string  `json:"product"`
	Quantity   int64   `json:"quantity"`
	TotalPrice float64 `json:"total_price"`
}
