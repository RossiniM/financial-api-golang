package module

import (
	"go.uber.org/fx"

	"github.com/go-chi/chi/v5"
)

func ProvideRouter() chi.Router {
	return chi.NewMux()
}

var Module = fx.Options(
	fx.Provide(ProvideRouter),
)
