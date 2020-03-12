package core_test

import (
	"testing"

	"github.com/carantes/superheroes-api/core"
	"github.com/stretchr/testify/suite"
)

type TestCoreUtilsSuite struct {
	suite.Suite
}

func (s *TestCoreUtilsSuite) TestParseInt() {
	s.Equal(1234, core.ParseInt("1234"))
}

func TestUtilsSuite(t *testing.T) {
	suite.Run(t, new(TestCoreUtilsSuite))
}
