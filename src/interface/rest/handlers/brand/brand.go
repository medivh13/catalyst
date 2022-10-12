package brand_handlers

/*
 * Author      : Jody (jody.almaida@gmail.com)
 * Modifier    :
 * Domain      : catalyst
 */

import (
	"encoding/json"
	"net/http"

	dto "catalyst/src/app/dtos/brand"
	usecases "catalyst/src/app/use_cases/brand"
	common_error "catalyst/src/infra/errors"
	"catalyst/src/interface/rest/response"
)

type BrandHandlerInterface interface {
	CreateBrand(w http.ResponseWriter, r *http.Request)
}

type brandHandler struct {
	response response.IResponseClient
	usecase  usecases.BrandUCInterface
}

func NewBrandHandler(r response.IResponseClient, h usecases.BrandUCInterface) BrandHandlerInterface {
	return &brandHandler{
		response: r,
		usecase:  h,
	}
}

func (h *brandHandler) CreateBrand(w http.ResponseWriter, r *http.Request) {

	postDTO := dto.CreateBrandReqDTO{}
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

	err = h.usecase.CreateBrand(&postDTO)
	if err != nil {
		h.response.HttpError(w, common_error.NewError(common_error.FAILED_CREATE_DATA, err))
		return
	}

	h.response.JSON(
		w,
		"Successful Create New Brand",
		nil,
		nil,
	)
}
