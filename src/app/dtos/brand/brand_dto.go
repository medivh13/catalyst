package brand_dto

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

type BrandInterface interface {
	Validate() error
}

type CreateBrandReqDTO struct {
	Name string `json:"name"`
}

func (dto *CreateBrandReqDTO) Validate() error {
	if err := validation.ValidateStruct(
		dto,
		validation.Field(&dto.Name, validation.Required),
	); err != nil {
		return err
	}
	return nil
}
