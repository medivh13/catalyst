package route

import (
	brandHandler "catalyst/src/interface/rest/handlers/brand"
	orderHandler "catalyst/src/interface/rest/handlers/order"
	productHandler "catalyst/src/interface/rest/handlers/product"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func BrandAppRouter(bh brandHandler.BrandHandlerInterface) http.Handler {
	r := chi.NewRouter()

	r.Mount("/", BrandRouter(bh))

	return r
}

func ProductAppRouter(ph productHandler.ProductHandlerInterface) http.Handler {
	r := chi.NewRouter()

	r.Mount("/", ProductRouter(ph))

	return r
}

func OrderAppRouter(oh orderHandler.OrderHandlerInterface) http.Handler {
	r := chi.NewRouter()

	r.Mount("/", OrderRouter(oh))

	return r
}
