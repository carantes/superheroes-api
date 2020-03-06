package superheroesbundle_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/carantes/superheroes-api/bundles/superheroesbundle"
	"github.com/stretchr/testify/assert"
)

func TestSearchSuperhero(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		assert.Equal(t, "/search/batman", req.URL.String())
		rw.Write([]byte(`{"Results": [{"Name": "Batman"}]}`))
	}))
	defer server.Close()

	// Get SuperheroAPI instance & call search
	api := superheroesbundle.NewSuperheroAPI(server.Client(), server.URL)
	response, err := api.Search("batman")
	assert.Nil(t, err, "should not return error on search")
	assert.Equal(t, "Batman", response.Results[0].Name)
}
