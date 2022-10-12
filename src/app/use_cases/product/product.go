package product_usecases

/*
 * Author      : Jody (jody.almaida@gmail.com)
 * Modifier    :
 * Domain      : catalyst
 */

import (
	dto "catalyst/src/app/dtos/product"
	"catalyst/src/domain/repositories"
	"log"
)

type ProductUCInterface interface {
	CreateProduct(data *dto.CreateProductReqDTO) error
	GetSingleProduct(data *dto.GetProductReqDTO) (*dto.GetSingleProductRespDTO, error)
	GetProductByBrandID(data *dto.GetProductReqDTO) ([]*dto.GetSingleProductRespDTO, error)
}

type productUseCase struct {
	ProductRepo repositories.ProductRepository
}

func NewProductUseCase(productRepo repositories.ProductRepository) *productUseCase {
	return &productUseCase{
		ProductRepo: productRepo,
	}
}

func (uc *productUseCase) CreateProduct(data *dto.CreateProductReqDTO) error {

	err := uc.ProductRepo.CreateProduct(data)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (uc *productUseCase) GetSingleProduct(data *dto.GetProductReqDTO) (*dto.GetSingleProductRespDTO, error) {

	result, err := uc.ProductRepo.GetSingleProduct(data)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return dto.ToReturnProduct(result), nil
}

func (uc *productUseCase) GetProductByBrandID(data *dto.GetProductReqDTO) ([]*dto.GetSingleProductRespDTO, error) {

	result, err := uc.ProductRepo.GetProductByBrandId(data)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return dto.ToProducts(result), nil
}
