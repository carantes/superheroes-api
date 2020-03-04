package superheroesbundle_test

import (
	"testing"

	"github.com/carantes/superheroes-api/bundles/superheroesbundle"
)

func TestNewSuperheroSpec(t *testing.T) {
	sh := superheroesbundle.NewSuperhero("Batman", "Bruce Wayne", 100, 47, "-")

	t.Run("superhero should have Name", func(t *testing.T) {
		expected := "Batman"
		if sh.Name != expected {
			t.Errorf("TestNewSuperheroSpec result is incorrect, return %s, expected %s", sh.Name, expected)
		}
	})

	t.Run("superhero should have FullName", func(t *testing.T) {
		expected := "Bruce Wayne"
		if sh.FullName != expected {
			t.Errorf("TestNewSuperheroSpec result is incorrect, return %s, expected %s", sh.FullName, expected)
		}
	})

	t.Run("superhero should have Intelligence", func(t *testing.T) {
		expected := 100
		if sh.Intelligence != expected {
			t.Errorf("TestNewSuperheroSpec result is incorrect, return %d, expected %d", sh.Intelligence, expected)
		}
	})

	t.Run("superhero should have Power", func(t *testing.T) {
		expected := 47
		if sh.Power != expected {
			t.Errorf("TestNewSuperheroSpec result is incorrect, return %d, expected %d", sh.Power, expected)
		}
	})

	t.Run("superhero should have Ocuppation", func(t *testing.T) {
		expected := "-"
		if sh.Occupation != expected {
			t.Errorf("TestNewSuperheroSpec result is incorrect, return %s, expected %s", sh.Occupation, expected)
		}
	})
}
