package superheroesbundle_test

import (
	"testing"

	"github.com/carantes/superheroes-api/bundles/superheroesbundle"
	"github.com/stretchr/testify/assert"
)

func TestNewSuperheroSpec(t *testing.T) {
	sh := superheroesbundle.NewSuperhero("Batman", "Bruce Wayne", 100, 47, "-")

	t.Run("superhero should have Name", func(t *testing.T) {
		assert.Equal(t, "Batman", sh.Name)
	})

	t.Run("superhero should have FullName", func(t *testing.T) {
		assert.Equal(t, "Bruce Wayne", sh.FullName)
	})

	t.Run("superhero should have Intelligence", func(t *testing.T) {
		assert.Equal(t, 100, sh.Intelligence)
	})

	t.Run("superhero should have Power", func(t *testing.T) {
		assert.Equal(t, 47, sh.Power)
	})

	t.Run("superhero should have Ocuppation", func(t *testing.T) {
		assert.Equal(t, "-", sh.Occupation)
	})
}
