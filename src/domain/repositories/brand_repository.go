package repositories

/*
 * Author      : Jody (jody.almaida@gmail.com)
 * Modifier    :
 * Domain      : catalyst
 */

import (
	dto "catalyst/src/app/dtos/brand"
)

type BrandRepository interface {
	CreateBrand(dataBrand *dto.CreateBrandReqDTO) error
}
