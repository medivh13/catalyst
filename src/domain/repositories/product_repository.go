package repositories

import (
	dto "catalyst/src/app/dtos/product"
	"catalyst/src/infra/models"
)

/*
 * Author      : Jody (jody.almaida@gmail.com)
 * Modifier    :
 * Domain      : catalyst
 */

// models "catalyst/src/infra/models"

type ProductRepository interface {
	CreateProduct(dataProduct *dto.CreateProductReqDTO) error
	GetSingleProduct(dataProduct *dto.GetProductReqDTO) (*models.Products, error)
	GetProductByBrandId(dataProduct *dto.GetProductReqDTO) ([]*models.Products, error)
}
