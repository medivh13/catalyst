package brand_usecases

/*
 * Author      : Jody (jody.almaida@gmail.com)
 * Modifier    :
 * Domain      : catalyst
 */

import (
	dto "catalyst/src/app/dtos/brand"
	"catalyst/src/domain/repositories"
	"log"
)

type BrandUCInterface interface {
	CreateBrand(data *dto.CreateBrandReqDTO) error
}

type brandUseCase struct {
	BrandRepo repositories.BrandRepository
}

func NewBrandUseCase(brandRepo repositories.BrandRepository) *brandUseCase {
	return &brandUseCase{
		BrandRepo: brandRepo,
	}
}

func (uc *brandUseCase) CreateBrand(data *dto.CreateBrandReqDTO) error {

	err := uc.BrandRepo.CreateBrand(data)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
