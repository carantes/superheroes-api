package superheroesbundle

import (
	"net/http"

	"github.com/carantes/superheroes-api/core"
)

// SuperheroesBundle handle superhero resources
type SuperheroesBundle struct {
	routes []core.Route
}

// NewSuperheroesBundle instance
func NewSuperheroesBundle() core.Bundle {
	ctr := NewSuperheroesController()

	// Bundle routes
	r := []core.Route{
		core.Route{
			Method:  http.MethodGet,
			Path:    "/superheroes",
			Handler: ctr.Index,
		},
	}

	return &SuperheroesBundle{
		routes: r,
	}
}

// GetRoutes implement core.Bundle interface
func (b *SuperheroesBundle) GetRoutes() []core.Route {
	return b.routes
}
