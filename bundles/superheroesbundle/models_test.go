package superheroesbundle_test

import (
	"testing"

	"github.com/carantes/superheroes-api/bundles/superheroesbundle"
	"github.com/stretchr/testify/suite"
)

type TestSuperheroesModelSuite struct {
	suite.Suite
	sh *superheroesbundle.Superhero
}

func (s *TestSuperheroesModelSuite) SetupSuite() {
	s.sh = superheroesbundle.NewSuperhero("Batman", "Bruce Wayne", 100, 47, "-")
}

func (s *TestSuperheroesModelSuite) TestNewSuperhero() {

	s.Run("superhero should have Name", func() {
		s.Equal("Batman", s.sh.Name)
	})

	s.Run("superhero should have FullName", func() {
		s.Equal("Bruce Wayne", s.sh.FullName)
	})

	s.Run("superhero should have Intelligence", func() {
		s.Equal(100, s.sh.Intelligence)
	})

	s.Run("superhero should have Power", func() {
		s.Equal(47, s.sh.Power)
	})

	s.Run("superhero should have Ocuppation", func() {
		s.Equal("-", s.sh.Occupation)
	})
}

func TestModelSuite(t *testing.T) {
	suite.Run(t, new(TestSuperheroesModelSuite))
}
