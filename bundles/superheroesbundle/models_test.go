package superheroesbundle_test

import (
	"testing"

	"github.com/carantes/superheroes-api/bundles/superheroesbundle"
	"github.com/stretchr/testify/assert"
)

func TestNewSuperheroSpec(t *testing.T) {
	sh := superheroesbundle.NewSuperhero("Batman", "Bruce Wayne", 100, 47, "-")

	t.Run("superhero should have Name", func(t *testing.T) {
		assert.Equal(t, sh.Name, "Batman")
	})

	t.Run("superhero should have FullName", func(t *testing.T) {
		assert.Equal(t, sh.FullName, "Bruce Wayne")
	})

	t.Run("superhero should have Intelligence", func(t *testing.T) {
		assert.Equal(t, sh.Intelligence, 100)
	})

	t.Run("superhero should have Power", func(t *testing.T) {
		assert.Equal(t, sh.Power, 47)
	})

	t.Run("superhero should have Ocuppation", func(t *testing.T) {
		assert.Equal(t, sh.Occupation, "-")
	})
}
