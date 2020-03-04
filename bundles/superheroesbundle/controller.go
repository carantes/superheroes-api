package superheroesbundle

import (
	"io"
	"net/http"
)

// SuperheroesController handle controller methods
type SuperheroesController struct {
}

// NewSuperheroesController instance
func NewSuperheroesController() *SuperheroesController {
	return &SuperheroesController{}
}

// Index return all superheroes
func (c *SuperheroesController) Index(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, `{"message": "Superheroes are awesome!"}`)
}
