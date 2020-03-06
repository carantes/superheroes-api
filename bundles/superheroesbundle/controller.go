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
	FindAll() ([]Superhero, error)
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
	log.Println("Get all superheroes")
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
	log.Printf("Get superhero id %s", uuid)

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
	if !sh.Validate() {
		c.HandleError(w, core.NewHTTPError(err, http.StatusBadRequest, "Bad request. Invalid payload"))
		return
	}

	// TODO: Check if superhero exist on database

	// Search for superheroes
	if c.service != nil {
		log.Println("Search for superhero on API")
		res, err := c.service.Search(sh.Name)

		if c.HandleError(w, err) {
			return
		}

		if len(res.Results) == 0 {
			c.HandleError(w, core.NewHTTPError(nil, http.StatusNotFound, "Not found. Can not found a superhero with this name"))
			return
		}

		//TODO: Handle existent and duplicates returns
		//TODO: Refactoring, add other fields
		sh.FullName = res.Results[0].Biography.FullName
		sh.Intelligence = core.GetUtils().ParseInt(res.Results[0].Powerstats.Intelligence)
		sh.Power = core.GetUtils().ParseInt(res.Results[0].Powerstats.Power)
		sh.Occupation = res.Results[0].Work.Occupation
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
	log.Printf("Delete superhero id %s", uuid)

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
