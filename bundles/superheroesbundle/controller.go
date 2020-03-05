package superheroesbundle

import (
	"log"
	"net/http"

	"github.com/carantes/superheroes-api/core"
	"github.com/gorilla/mux"
	uuid "github.com/satori/go.uuid"
)

// SuperheroesController handle controller methods
type SuperheroesController struct {
	core.Controller
	repo SuperheroesRepositoryInterface
}

// SuperheroesRepositoryInterface define the interface to retrieve, insert and delete data from repository
type SuperheroesRepositoryInterface interface {
	FindAll() ([]Superhero, error)
	FindOne(uuid.UUID) (Superhero, error)
	Insert(*Superhero) error
	Delete(id uuid.UUID) error
}

// NewSuperheroesController instance
func NewSuperheroesController(repo SuperheroesRepositoryInterface) *SuperheroesController {
	return &SuperheroesController{
		repo: repo,
	}
}

// Index return all superheroes
func (c *SuperheroesController) Index(w http.ResponseWriter, r *http.Request) {
	sh, err := c.repo.FindAll()

	if c.HandleError(w, err) {
		return
	}

	c.SendJSON(w, &sh, http.StatusOK)
}

// Get return one superhero by UUID
func (c *SuperheroesController) Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uuid, err := uuid.FromString(vars["uuid"])
	log.Println(uuid, err)

	if err != nil {
		c.HandleError(w, core.NewHTTPError(err, http.StatusBadRequest, "Bad request. Invalid UUID."))
		return
	}

	sh, err := c.repo.FindOne(uuid)
	if c.HandleError(w, err) {
		return
	}

	c.SendJSON(w, &sh, http.StatusOK)
}

// Create a new superhero
func (c *SuperheroesController) Create(w http.ResponseWriter, r *http.Request) {
	var sh Superhero

	// Parse request body to JSON
	err := c.GetContent(r, &sh)

	if err != nil {
		c.HandleError(w, core.NewHTTPError(err, http.StatusBadRequest, "Bad request. Invalid JSON"))
		return
	}

	// Validate superhero data
	if !sh.Validate() {
		c.HandleError(w, core.NewHTTPError(err, http.StatusBadRequest, "Bad request. Invalid payload"))
		return
	}

	// Insert superhero on repository
	err = c.repo.Insert(&sh)

	if c.HandleError(w, err) {
		return
	}

	c.SendJSON(w, &sh, http.StatusCreated)
}

// Delete a superhero
func (c *SuperheroesController) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uuid, err := uuid.FromString(vars["uuid"])

	if err != nil {
		c.HandleError(w, core.NewHTTPError(err, http.StatusBadRequest, "Bad request. Invalid UUID."))
		return
	}

	// Delete superhero from repository
	err = c.repo.Delete(uuid)

	if c.HandleError(w, err) {
		return
	}

	c.SendJSON(w, nil, http.StatusNoContent)
}
