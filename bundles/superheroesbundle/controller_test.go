package superheroesbundle_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/carantes/superheroes-api/bundles/superheroesbundle"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func GetTestController(data []superheroesbundle.Superhero) *superheroesbundle.SuperheroesController {
	r := superheroesbundle.NewSuperheroesRepository(data)
	c := superheroesbundle.NewSuperheroesController(r, nil)
	return c
}

func TestSuperheroesControllerIndexSpec(t *testing.T) {
	c := GetTestController([]superheroesbundle.Superhero{
		*superheroesbundle.NewSuperhero("Batman", "Bruce Wayne", 100, 47, "-"),
		*superheroesbundle.NewSuperhero("Wolverine", "Logan", 63, 89, "Adventurer, instructor, former bartender..."),
	})

	r := mux.NewRouter()
	r.HandleFunc("/superheroes", c.Index).Methods("GET")

	t.Run("GET all superheroes", func(t *testing.T) {
		rr := httptest.NewRecorder()
		req, err := http.NewRequest("GET", "/superheroes", nil)
		assert.Nil(t, err, "should not return error on new request")
		r.ServeHTTP(rr, req)

		var sh []superheroesbundle.Superhero
		err = json.Unmarshal([]byte(rr.Body.String()), &sh)
		assert.Nil(t, err, "should not return error on decode to a list of superheroes")

		assert.Equal(t, http.StatusOK, rr.Code, "should return status 200 OK")
		assert.Equal(t, 2, len(sh), "should return two superheroes")
		assert.Equal(t, "Batman", sh[0].Name, "first superhero should be Batman")
	})
}

func TestSuperheroesControllerGetSpec(t *testing.T) {
	sh := superheroesbundle.NewSuperhero("Batman", "Bruce Wayne", 100, 47, "-")
	c := GetTestController([]superheroesbundle.Superhero{*sh})

	r := mux.NewRouter()
	r.HandleFunc("/superheroes/{uuid}", c.Get).Methods("GET")

	t.Run("GET superhero with an invalid UUID", func(t *testing.T) {
		rr := httptest.NewRecorder()
		req, err := http.NewRequest("GET", "/superheroes/1234", nil)
		assert.Nil(t, err, "should not return error on new request")
		r.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusBadRequest, rr.Code, "should return status 400 Bad Request")
		assert.Equal(t, `{"detail":"Bad request. Invalid UUID."}`, rr.Body.String(), "should return empty body")
	})

	t.Run("GET superhero with a valid UUUID", func(t *testing.T) {
		rr := httptest.NewRecorder()
		req, err := http.NewRequest("GET", "/superheroes/"+sh.ID.String(), nil)
		assert.Nil(t, err, "should not return error on new request")
		r.ServeHTTP(rr, req)

		var sh superheroesbundle.Superhero
		err = json.Unmarshal([]byte(rr.Body.String()), &sh)
		assert.Nil(t, err, "should not return error on decode to a list of superheroes")

		assert.Equal(t, http.StatusOK, rr.Code, "should return status 200 OK")
		assert.Equal(t, "Batman", sh.Name, "first superhero should be Batman")
	})

}

func TestSuperheroesControllerCreateSpec(t *testing.T) {
	c := GetTestController([]superheroesbundle.Superhero{})

	r := mux.NewRouter()
	r.HandleFunc("/superheroes", c.Create).Methods("POST")

	t.Run("POST empty payload to create superhero return bad request", func(t *testing.T) {
		rr := httptest.NewRecorder()
		req, err := http.NewRequest("POST", "/superheroes", bytes.NewBufferString(``))
		assert.Nil(t, err, "should not return error on new request")
		r.ServeHTTP(rr, req)

		assert.Equal(t, 400, rr.Code, "should return status 400 bad request")
		assert.Equal(t, `{"detail":"Bad request. Invalid JSON"}`, rr.Body.String(), "should return invalid json message")
	})

	t.Run("POST invalid payload to create superhero return bad request", func(t *testing.T) {
		rr := httptest.NewRecorder()
		req, err := http.NewRequest("POST", "/superheroes", bytes.NewBufferString(`{ "Name": "" }`))
		assert.Nil(t, err, "should not return error on new request")
		r.ServeHTTP(rr, req)

		assert.Equal(t, 400, rr.Code, "should return status 400 bad request")
		assert.Equal(t, `{"detail":"Bad request. Invalid payload"}`, rr.Body.String(), "should return invalid payload message")
	})

	t.Run("POST valid payload return superhero and status 201 created", func(t *testing.T) {
		rr := httptest.NewRecorder()
		req, err := http.NewRequest("POST", "/superheroes", bytes.NewBufferString(`{ "Name": "Batman" }`))
		assert.Nil(t, err, "should not return error on new request")
		r.ServeHTTP(rr, req)

		var sh superheroesbundle.Superhero
		err = json.Unmarshal([]byte(rr.Body.String()), &sh)
		assert.Nil(t, err, "should not return error on decode a superhero")
		assert.Equal(t, rr.Code, 201, "should return status 201 created")
		assert.Equal(t, sh.Name, "Batman")
	})
}

func TestSuperheroesControllerDeleteSpec(t *testing.T) {
	sh1 := superheroesbundle.NewSuperhero("Batman", "Bruce Wayne", 100, 47, "-")
	sh2 := superheroesbundle.NewSuperhero("Wolverine", "Logan", 63, 89, "Adventurer, instructor, former bartender...")
	c := GetTestController([]superheroesbundle.Superhero{*sh1, *sh2})

	r := mux.NewRouter()
	r.HandleFunc("/superheroes/{uuid}", c.Delete).Methods("DELETE")

	t.Run("DELETE superhero with an invalid UUID", func(t *testing.T) {
		rr := httptest.NewRecorder()
		req, err := http.NewRequest("DELETE", "/superheroes/1234", nil)
		assert.Nil(t, err, "should not return error on new request")
		r.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusBadRequest, rr.Code, "should return status 400 Bad Request")
		assert.Equal(t, `{"detail":"Bad request. Invalid UUID."}`, rr.Body.String(), "should return empty body")
	})

	t.Run("DELETE superhero with a valid UUID", func(t *testing.T) {
		rr := httptest.NewRecorder()
		req, err := http.NewRequest("DELETE", "/superheroes/"+sh2.ID.String(), nil)
		assert.Nil(t, err, "should not return error on new request")
		r.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusNoContent, rr.Code, "should return status 204 No Content")
	})
}
