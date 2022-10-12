package route

import (
	"net/http"

	brandHandler "catalyst/src/interface/rest/handlers/brand"

	"github.com/go-chi/chi/v5"
)

func BrandRouter(h brandHandler.BrandHandlerInterface) http.Handler {
	r := chi.NewRouter()

	r.Post("/", h.CreateBrand)
	return r
}
