package usecases

import (
	brandUC "catalyst/src/app/use_cases/brand"
	orderUC "catalyst/src/app/use_cases/order"
	productUC "catalyst/src/app/use_cases/product"
)

type AllUseCases struct {
	BrandUseCase   brandUC.BrandUCInterface
	ProductUseCase productUC.ProductUCInterface
	OrderUseCase   orderUC.OrderUCInterface
}
