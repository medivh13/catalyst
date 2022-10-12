package products_dto

import (
	models "catalyst/src/infra/models"
)

func ToProducts(datas []*models.Products) []*GetSingleProductRespDTO {
	var resp []*GetSingleProductRespDTO
	for _, m := range datas {
		resp = append(resp, ToReturnProduct(m))
	}
	return resp
}

func ToReturnProduct(d *models.Products) *GetSingleProductRespDTO {
	return &GetSingleProductRespDTO{
		ID:      d.ID,
		Name:    d.Name,
		Price:   d.Price,
		BrandId: d.BrandID,
	}
}
