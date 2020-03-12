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
	repo    SuperheroesRepositoryInterface
	service SuperheroesServiceInterface
}

// SuperheroesRepositoryInterface define the interface to retrieve, insert and delete data from repository
type SuperheroesRepositoryInterface interface {
	FindAll(*Superhero) ([]Superhero, error)
	FindOne(uuid.UUID) (Superhero, error)
	Insert(*Superhero) error
	Delete(id uuid.UUID) error
}

// SuperheroesServiceInterface define the interface with external superheroapi service
type SuperheroesServiceInterface interface {
	Search(string) (SuperheroAPISearchResponse, error)
}

// NewSuperheroesController instance
func NewSuperheroesController(repo SuperheroesRepositoryInterface, service SuperheroesServiceInterface) *SuperheroesController {
	return &SuperheroesController{
		repo:    repo,
		service: service,
	}
}

// Index return all superheroes
func (c *SuperheroesController) Index(w http.ResponseWriter, r *http.Request) {
	var a SuperheroAlignment

	query := r.URL.Query()

	// Filters
	filter := Superhero{
		Name:      query.Get("name"),
		Alignment: a.FromString(query.Get("alignment")),
	}

	sh, err := c.repo.FindAll(&filter)

	if c.HandleError(w, err) {
		return
	}

	c.SendJSON(w, &sh, http.StatusOK)
}

// Get return one superhero by UUID
func (c *SuperheroesController) Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uuid, err := uuid.FromString(vars["uuid"])

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
	log.Println("Creating superhero")

	// Parse request body to JSON
	err := c.GetContent(r, &sh)

	if err != nil {
		c.HandleError(w, core.NewHTTPError(err, http.StatusBadRequest, "Bad request. Invalid JSON"))
		return
	}

	// Validate superhero data
	if err := sh.Validate(); err != nil {
		c.HandleError(w, core.NewHTTPError(err, http.StatusBadRequest, "Bad request. Invalid payload"))
		return
	}

	// Search superhero on local repository
	current, err := c.repo.FindAll(&sh)

	// Return the current database version and response 200 (OK)
	if err == nil && len(current) > 0 {
		c.SendJSON(w, &current[0], http.StatusOK)
		return
	}

	// Search for it on superheroapi.com
	res, err := c.service.Search(sh.Name)

	if c.HandleError(w, err) {
		return
	}

	if len(res.Results) == 0 {
		c.HandleError(w, core.NewHTTPError(nil, http.StatusNotFound, "Not found. Can not found a superhero with this name"))
		return
	}

	// Get the first match and mapping to superhero object
	sh = res.ToSuperhero()[0]

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
