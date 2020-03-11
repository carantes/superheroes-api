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
	s.sh = superheroesbundle.NewSuperhero("Batman", "Bruce Wayne", superheroesbundle.GoodAlignment, 100, 47, "-", "http://www.hero.com/10441.jpg", 10)
}

func (s *TestSuperheroesModelSuite) TestSuperheroAlignment() {
	s.Run("superhero alignment should return a string", func() {
		good := superheroesbundle.GoodAlignment
		s.Equal("good", good.ToString())
	})

	s.Run("superhero alignment should be created from a string", func() {
		var good superheroesbundle.SuperheroAlignment
		s.Equal(superheroesbundle.GoodAlignment, good.FromString("good"))
	})
}

func (s *TestSuperheroesModelSuite) TestNewSuperhero() {
	s.Equal("Batman", s.sh.Name, "superhero should have Name")
	s.Equal("Bruce Wayne", s.sh.FullName, "superhero should have FullName")
	s.Equal(superheroesbundle.GoodAlignment, s.sh.Alignment, "superhero should have alignment")
	s.Equal(100, s.sh.Intelligence, "superhero should have Intelligence")
	s.Equal(47, s.sh.Power, "superhero should have Power")
	s.Equal("-", s.sh.Occupation, "superhero should have Ocuppation")
	s.Equal("http://www.hero.com/10441.jpg", s.sh.ImageURL, "superhero should have Image")
	s.Equal(10, s.sh.TotalRelatives, "superhero should have Relatives")
}

func (s *TestSuperheroesModelSuite) TestAddGroup() {
	sh := superheroesbundle.NewSuperhero("Batman", "Bruce Wayne", superheroesbundle.GoodAlignment, 100, 47, "-", "http://www.hero.com/10441.jpg", 10)
	s.Len(sh.Groups, 0)
	sh.AddGroup("Justice League")
	s.Len(sh.Groups, 1)
}

func TestModelSuite(t *testing.T) {
	suite.Run(t, new(TestSuperheroesModelSuite))
}
