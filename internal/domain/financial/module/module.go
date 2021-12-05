package module

import (
	"example.com/m/v2/internal/domain/financial/endpoint"
	"example.com/m/v2/internal/domain/financial/service"
	"github.com/go-chi/chi/v5"
	"go.uber.org/fx"
)

//interface

func ProvideService() *service.Service {
	return &service.Service{Message: "Hello", Destination: "World"}
}

func bindEndpoint(router chi.Router,
	service *service.Service) {
	endpoint.Bind(router, service)
}

// Module provided to fx
var Module = fx.Options(
	fx.Provide(ProvideService),
	fx.Invoke(bindEndpoint),
)
