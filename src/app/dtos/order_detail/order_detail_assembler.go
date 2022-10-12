package order_detail_dto

import "catalyst/src/infra/models"

func ToReturnOrder(d []*models.Orders) *GetOrderRespDTO {
	resp := &GetOrderRespDTO{
		ID:        d[0].ID,
		OrderCode: d[0].OrderCode,
		Total:     d[0].Total,
		Detail:    ToOrderDetails(d),
	}

	return resp
}

func ToOrderDetails(datas []*models.Orders) []*GetOrderDetailRespDTO {
	var resp []*GetOrderDetailRespDTO
	for _, m := range datas {
		resp = append(resp, ToReturnOrderDetail(m))
	}
	return resp
}

func ToReturnOrderDetail(d *models.Orders) *GetOrderDetailRespDTO {
	return &GetOrderDetailRespDTO{
		Product:    d.Product,
		Quantity:   d.Quantity,
		TotalPrice: d.TotalPrice,
	}
}
