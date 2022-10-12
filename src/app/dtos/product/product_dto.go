package products_dto

import validation "github.com/go-ozzo/ozzo-validation"

type PrroductInterface interface {
	Validate() error
}

type CreateProductReqDTO struct {
	Name    string  `json:"name"`
	Price   float64 `json:"price"`
	BrandId int64   `json:"brand_id"`
}

func (dto *GetProductReqDTO) Validate() error {
	if err := validation.ValidateStruct(
		dto,
		validation.Field(&dto.ID, validation.Required),
	); err != nil {
		return err
	}
	return nil
}

type GetProductReqDTO struct {
	ID int64 `json:"id"`
}

func (dto *CreateProductReqDTO) Validate() error {
	if err := validation.ValidateStruct(
		dto,
		validation.Field(&dto.Name, validation.Required),
		validation.Field(&dto.Price, validation.Required),
		validation.Field(&dto.BrandId, validation.Required),
	); err != nil {
		return err
	}
	return nil
}

type GetSingleProductRespDTO struct {
	ID      int64   `json:"id"`
	Name    string  `json:"name"`
	Price   float64 `json:"price"`
	BrandId int64   `json:"brand_id"`
}
