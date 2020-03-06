package superheroesbundle

import (
	"net/http"
	"time"

	"github.com/carantes/superheroes-api/core"
)

// SuperheroesBundle handle superhero resources
type SuperheroesBundle struct {
	routes []core.Route
}

// NewSuperheroesBundle instance
func NewSuperheroesBundle() core.Bundle {
	data := []Superhero{
		*NewSuperhero("Batman", "Bruce Wayne", 100, 47, "-"),
		*NewSuperhero("Wolverine", "Logan", 63, 89, "Adventurer, instructor, former bartender..."),
	}

	repo := NewSuperheroesRepository(data)
	cfg := core.GetConfig()
	httpClient := &http.Client{Timeout: time.Second * time.Duration(cfg.SuperheroAPITimeout)}
	baseURL := cfg.SuperheroAPIBaseURL + cfg.SuperheroAPIToken
	api := NewSuperheroAPI(httpClient, baseURL)
	ctr := NewSuperheroesController(repo, api)

	// Bundle routes
	r := []core.Route{
		core.Route{
			Method:  http.MethodGet,
			Path:    "/superheroes",
			Handler: ctr.Index,
		},
		core.Route{
			Method:  http.MethodGet,
			Path:    "/superheroes/{uuid}",
			Handler: ctr.Get,
		},
		core.Route{
			Method:  http.MethodPost,
			Path:    "/superheroes",
			Handler: ctr.Create,
		},
		core.Route{
			Method:  http.MethodDelete,
			Path:    "/superheroes/{uuid}",
			Handler: ctr.Delete,
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
