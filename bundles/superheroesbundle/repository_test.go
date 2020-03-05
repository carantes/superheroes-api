package superheroesbundle_test

import (
	"testing"

	"github.com/carantes/superheroes-api/bundles/superheroesbundle"
	"github.com/stretchr/testify/assert"
)

func TestSuperheroesRepositorySpec(t *testing.T) {
	batman := superheroesbundle.NewSuperhero("Batman", "Bruce Wayne", 100, 47, "-")
	data := []superheroesbundle.Superhero{*batman}
	r := superheroesbundle.NewSuperheroesRepository(data)

	t.Run("Find all superheroes", func(t *testing.T) {
		sh, err := r.FindAll()
		assert.Nil(t, err, "should not return error")
		assert.Len(t, sh, 1)
	})

	t.Run("Find one superhero", func(t *testing.T) {
		sh, err := r.FindOne(batman.ID)
		assert.Nil(t, err, "should not return error")
		assert.Equal(t, "Batman", sh.Name)
	})

	t.Run("Create superhero", func(t *testing.T) {
		err := r.Insert(superheroesbundle.NewSuperhero("Wolverine", "Logan", 63, 89, "Adventurer, instructor, former bartender..."))
		assert.Nil(t, err, "should not return error")
	})

	t.Run("Create superhero", func(t *testing.T) {
		err := r.Delete(batman.ID)
		assert.Nil(t, err, "should not return error")
	})
}
