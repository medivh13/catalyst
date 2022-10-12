package route

import (
	"net/http"

	orderHandler "catalyst/src/interface/rest/handlers/order"

	"github.com/go-chi/chi/v5"
)

func OrderRouter(h orderHandler.OrderHandlerInterface) http.Handler {
	r := chi.NewRouter()

	r.Post("/", h.CreateOrder)
	r.Get("/{transactionid}", h.GetOrderByID)
	return r
}
