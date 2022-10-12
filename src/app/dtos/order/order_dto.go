package order_dto

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

type OrderInterface interface {
	Validate() error
}

type CreateOrderReqDTO struct {
	Data []*DetailOrderReqDTO `json:"data"`
}

type DetailOrderReqDTO struct {
	ProductID int64   `json:"product_id"`
	Price     float64 `json:"price"`
	Quantity  int64   `json:"qty"`
}

func (dto *CreateOrderReqDTO) Validate() error {
	if err := validation.ValidateStruct(
		dto,
		validation.Field(&dto.Data, validation.Required),
	); err != nil {
		return err
	}
	return nil
}

type OrderRespDTO struct {
	OrderID int64 `json:"id`
	// OrderCode string `json:"order_code`
}
