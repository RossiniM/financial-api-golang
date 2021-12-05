package endpoint

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

// type handler struct {
// 	router *chi.Mux
// }

type Category int

const (
	Home Category = iota
	Study
	Food
	Fuel
)

// type entry struct {
// 	value       float64
// 	description string
// 	dueDate     string
// 	category    Category
// }

type Service interface {
	Dispatch() string
}

type Endpoint struct {
	service Service
}

func Bind(
	router chi.Router,
	service Service,
) {
	endpoint := &Endpoint{service: service}
	router.Route("/api/transactions", func(r chi.Router) {
		r.Get("/", endpoint.get)
	})
}

func (endpoint *Endpoint) get(rw http.ResponseWriter, r *http.Request) {

	message := endpoint.service.Dispatch()
	rw.Write([]byte(message))
}

// id: number;
//     title: string;
//     amount: number;
//     type: string;
//     category: string;
//     createdAt: string;
