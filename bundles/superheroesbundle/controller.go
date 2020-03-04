package superheroesbundle

import (
	"net/http"

	"github.com/carantes/superheroes-api/core"
)

// SuperheroesController handle controller methods
type SuperheroesController struct {
	core.Controller
}

// NewSuperheroesController instance
func NewSuperheroesController() *SuperheroesController {
	return &SuperheroesController{}
}

// Index return all superheroes
func (c *SuperheroesController) Index(w http.ResponseWriter, r *http.Request) {
	sh := []*Superhero{
		NewSuperhero("Batman", "Bruce Wayne", 100, 47, "-"),
		NewSuperhero("Wolverine", "Logan", 63, 89, "Adventurer, instructor, former bartender, bouncer, spy, government operative, mercenary, soldier, sailor, miner"),
	}

	c.SendJSON(w, &sh, http.StatusOK)
}
