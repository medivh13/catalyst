package repositories

import (
	dto "catalyst/src/app/dtos/order"
	dtoDetail "catalyst/src/app/dtos/order_detail"
	"catalyst/src/infra/models"
)

/*
 * Author      : Jody (jody.almaida@gmail.com)
 * Modifier    :
 * Domain      : catalyst
 */

type OrderRepository interface {
	CreateOrder(dataOrder *dto.CreateOrderReqDTO) (*models.OrderReqModel, error)
	GetOrderByID(dataOrder *dtoDetail.GetOrderReqDTO) ([]*models.Orders, error)
}
