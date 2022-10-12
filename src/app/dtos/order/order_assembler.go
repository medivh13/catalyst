package order_dto

import (
	models "catalyst/src/infra/models"
)

func ToReturnOrderReq(d *models.OrderReqModel) *OrderRespDTO {
	return &OrderRespDTO{
		OrderID: d.OrderID,
		// OrderCode: d.OrderCode,
	}
}
