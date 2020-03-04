package superheroesbundle_test

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/carantes/superheroes-api/bundles/superheroesbundle"
	"github.com/stretchr/testify/assert"
)

func TestSuperheroesControllerSpec(t *testing.T) {
	c := superheroesbundle.NewSuperheroesController()
	mux := http.NewServeMux()
	server := httptest.NewServer(mux)
	defer server.Close()

	t.Run("GET all superheroes", func(t *testing.T) {
		mux.HandleFunc("/", c.Index)
		resp, err := http.Get(server.URL + "/")

		if err != nil {
			t.Fatal(err)
		}

		body, err := ioutil.ReadAll(resp.Body)

		assert.Nil(t, err, "should not return error")

		// TODO: validate body data against model
		assert.NotNil(t, body)
		assert.Equal(t, resp.StatusCode, http.StatusOK)
	})
}
