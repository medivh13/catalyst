package brand_handlers

/*
 * Author      : Jody (jody.almaida@gmail.com)
 * Modifier    :
 * Domain      : catalyst
 */

import (
	"encoding/json"
	"net/http"
	"strconv"

	dto "catalyst/src/app/dtos/product"
	usecases "catalyst/src/app/use_cases/product"
	common_error "catalyst/src/infra/errors"
	"catalyst/src/interface/rest/response"
)

type ProductHandlerInterface interface {
	CreateProduct(w http.ResponseWriter, r *http.Request)
	GetSingleProduct(w http.ResponseWriter, r *http.Request)
	GetProductByBrandID(w http.ResponseWriter, r *http.Request)
}

type productHandler struct {
	response response.IResponseClient
	usecase  usecases.ProductUCInterface
}

func NewProductHandler(r response.IResponseClient, h usecases.ProductUCInterface) ProductHandlerInterface {
	return &productHandler{
		response: r,
		usecase:  h,
	}
}

func (h *productHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {

	postDTO := dto.CreateProductReqDTO{}
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

	err = h.usecase.CreateProduct(&postDTO)
	if err != nil {
		h.response.HttpError(w, common_error.NewError(common_error.FAILED_CREATE_DATA, err))
		return
	}

	h.response.JSON(
		w,
		"Successful Create New Product",
		nil,
		nil,
	)
}

func (h *productHandler) GetSingleProduct(w http.ResponseWriter, r *http.Request) {

	getDTO := dto.GetProductReqDTO{}
	dt, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		h.response.HttpError(w, common_error.NewError(common_error.DATA_INVALID, err))
		return
	}
	getDTO.ID = int64(dt)

	err = getDTO.Validate()
	if err != nil {
		h.response.HttpError(w, common_error.NewError(common_error.DATA_INVALID, err))
		return
	}

	data, err := h.usecase.GetSingleProduct(&getDTO)
	if err != nil {
		h.response.HttpError(w, common_error.NewError(common_error.FAILED_RETRIEVE_DATA, err))
		return
	}

	h.response.JSON(
		w,
		"Successful Get Product",
		data,
		nil,
	)
}

func (h *productHandler) GetProductByBrandID(w http.ResponseWriter, r *http.Request) {

	getDTO := dto.GetProductReqDTO{}
	dt, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		h.response.HttpError(w, common_error.NewError(common_error.DATA_INVALID, err))
		return
	}
	getDTO.ID = int64(dt)

	err = getDTO.Validate()
	if err != nil {
		h.response.HttpError(w, common_error.NewError(common_error.DATA_INVALID, err))
		return
	}

	data, err := h.usecase.GetProductByBrandID(&getDTO)
	if err != nil {
		h.response.HttpError(w, common_error.NewError(common_error.FAILED_RETRIEVE_DATA, err))
		return
	}

	h.response.JSON(
		w,
		"Successful Get Product",
		data,
		nil,
	)
}
