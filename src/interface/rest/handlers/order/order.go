package order_handlers

/*
 * Author      : Jody (jody.almaida@gmail.com)
 * Modifier    :
 * Domain      : catalyst
 */

import (
	"encoding/json"
	"net/http"

	dto "catalyst/src/app/dtos/order"
	dtoDetail "catalyst/src/app/dtos/order_detail"
	usecases "catalyst/src/app/use_cases/order"
	common_error "catalyst/src/infra/errors"
	"catalyst/src/interface/rest/response"

	"github.com/go-chi/chi/v5"
)

type OrderHandlerInterface interface {
	CreateOrder(w http.ResponseWriter, r *http.Request)
	GetOrderByID(w http.ResponseWriter, r *http.Request)
}

type orderHandler struct {
	response response.IResponseClient
	usecase  usecases.OrderUCInterface
}

func NewOrderHandler(r response.IResponseClient, h usecases.OrderUCInterface) OrderHandlerInterface {
	return &orderHandler{
		response: r,
		usecase:  h,
	}
}

func (h *orderHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {

	postDTO := dto.CreateOrderReqDTO{}
	err := json.NewDecoder(r.Body).Decode(&postDTO)
	if err != nil {
		h.response.HttpError(w, common_error.NewError(common_error.DATA_INVALID, err))
		return
	}
	err = postDTO.Validate()
	if err != nil {
		h.response.HttpError(w, common_error.NewError(common_error.DATA_INVALID, err))
		return
	}

	data, err := h.usecase.CreateOrder(&postDTO)
	if err != nil {
		h.response.HttpError(w, common_error.NewError(common_error.FAILED_CREATE_DATA, err))
		return
	}

	h.response.JSON(
		w,
		"Successful Create Order",
		data,
		nil,
	)
}

func (h *orderHandler) GetOrderByID(w http.ResponseWriter, r *http.Request) {

	getDTO := dtoDetail.GetOrderReqDTO{}
	getDTO.OrderID = chi.URLParam(r, "transactionid")

	err := getDTO.Validate()
	if err != nil {
		h.response.HttpError(w, common_error.NewError(common_error.DATA_INVALID, err))
		return
	}

	data, err := h.usecase.GetOrderByID(&getDTO)
	if err != nil {
		h.response.HttpError(w, common_error.NewError(common_error.FAILED_CREATE_DATA, err))
		return
	}

	h.response.JSON(
		w,
		"Successful Get Order Data",
		data,
		nil,
	)
}
