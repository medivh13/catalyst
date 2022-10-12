package route

import (
	"net/http"

	productHandler "catalyst/src/interface/rest/handlers/product"

	"github.com/go-chi/chi/v5"
)

func ProductRouter(h productHandler.ProductHandlerInterface) http.Handler {
	r := chi.NewRouter()

	r.Post("/", h.CreateProduct)
	r.Get("/", h.GetSingleProduct)
	r.Get("/brand", h.GetProductByBrandID)
	return r
}
