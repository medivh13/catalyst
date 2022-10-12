package order_usecases

/*
 * Author      : Jody (jody.almaida@gmail.com)
 * Modifier    :
 * Domain      : catalyst
 */

import (
	dto "catalyst/src/app/dtos/order"
	dtoDetail "catalyst/src/app/dtos/order_detail"
	"catalyst/src/domain/repositories"
	"log"
)

type OrderUCInterface interface {
	CreateOrder(data *dto.CreateOrderReqDTO) (*dto.OrderRespDTO, error)
	GetOrderByID(data *dtoDetail.GetOrderReqDTO) (*dtoDetail.GetOrderRespDTO, error)
}

type orderUseCase struct {
	OrderRepo repositories.OrderRepository
}

func NewOrderUseCase(orderRepo repositories.OrderRepository) *orderUseCase {
	return &orderUseCase{
		OrderRepo: orderRepo,
	}
}

func (uc *orderUseCase) CreateOrder(data *dto.CreateOrderReqDTO) (*dto.OrderRespDTO, error) {

	result, err := uc.OrderRepo.CreateOrder(data)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return dto.ToReturnOrderReq(result), nil
}

func (uc *orderUseCase) GetOrderByID(data *dtoDetail.GetOrderReqDTO) (*dtoDetail.GetOrderRespDTO, error) {

	result, err := uc.OrderRepo.GetOrderByID(data)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return dtoDetail.ToReturnOrder(result), nil
}
