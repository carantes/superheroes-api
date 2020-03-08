package superheroesbundle_test

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/carantes/superheroes-api/bundles/superheroesbundle"
	"github.com/gorilla/mux"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type TestSuperheroesControllerSuite struct {
	suite.Suite
	router     *mux.Router
	controller *superheroesbundle.SuperheroesController
	repo       *SuperheroRepositoryMock
}

// Repository Mock
// TODO: migrate this to a helper file
type SuperheroRepositoryMock struct {
	mock.Mock
}

func (mock *SuperheroRepositoryMock) FindAll() ([]superheroesbundle.Superhero, error) {
	args := mock.Called()
	return args.Get(0).([]superheroesbundle.Superhero), args.Error(1)
}

func (mock *SuperheroRepositoryMock) FindOne(id uuid.UUID) (superheroesbundle.Superhero, error) {
	args := mock.Called(id)
	return args.Get(0).(superheroesbundle.Superhero), args.Error(1)
}

func (mock *SuperheroRepositoryMock) Insert(sh *superheroesbundle.Superhero) error {
	args := mock.Called(sh)
	return args.Error(0)
}

func (mock *SuperheroRepositoryMock) Delete(id uuid.UUID) error {
	args := mock.Called(id)
	return args.Error(0)
}

func (s *TestSuperheroesControllerSuite) SetupSuite() {
	s.repo = new(SuperheroRepositoryMock)
	s.controller = superheroesbundle.NewSuperheroesController(s.repo, nil)
	s.router = mux.NewRouter()
}

func (s *TestSuperheroesControllerSuite) ServeRequest(method string, url string, payload io.Reader) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	req, err := http.NewRequest(method, url, payload)
	s.NoError(err, "should not return error on new request")
	s.router.ServeHTTP(rr, req)
	return rr
}

func (s *TestSuperheroesControllerSuite) TestIndex() {
	// Mock repository FindAll function to return mocked data
	mockData := []superheroesbundle.Superhero{
		*superheroesbundle.NewSuperhero("Batman", "Bruce Wayne", 100, 47, "-"),
		*superheroesbundle.NewSuperhero("Wolverine", "Logan", 63, 89, "Adventurer, instructor, former bartender..."),
	}
	s.repo.On("FindAll").Return(mockData, nil)

	// Setup handler for Index function
	s.router.HandleFunc("/superheroes", s.controller.Index).Methods("GET")

	// Call index
	rr := s.ServeRequest("GET", "/superheroes", nil)

	// Assert return expected result
	s.Equal(200, rr.Code, "should return status 200 success")
	s.repo.AssertExpectations(s.T())
}

func (s *TestSuperheroesControllerSuite) TestGetInvalidUUID() {
	s.router.HandleFunc("/superheroes/{uuid}", s.controller.Get).Methods("GET")

	rr := s.ServeRequest("GET", "/superheroes/1234", nil)

	s.Equal(http.StatusBadRequest, rr.Code, "should return status 400 Bad Request")
	s.Equal(`{"detail":"Bad request. Invalid UUID."}`, rr.Body.String(), "should return invalid UUID message")
}

func (s *TestSuperheroesControllerSuite) TestGet() {
	s.router.HandleFunc("/superheroes/{uuid}", s.controller.Get).Methods("GET")

	mockData := *superheroesbundle.NewSuperhero("Batman", "Bruce Wayne", 100, 47, "-")
	s.repo.On("FindOne", mockData.ID).Return(mockData, nil)

	rr := s.ServeRequest("GET", "/superheroes/"+mockData.ID.String(), nil)

	// Assert return expected result
	s.Equal(200, rr.Code, "should return status 200 success")
	s.repo.AssertExpectations(s.T())
}

func (s *TestSuperheroesControllerSuite) TestCreateEmptyPayload() {
	s.router.HandleFunc("/superheroes", s.controller.Create).Methods("POST")

	rr := s.ServeRequest("POST", "/superheroes", bytes.NewBufferString(``))

	s.Equal(400, rr.Code, "should return status 400 bad request")
	s.Equal(`{"detail":"Bad request. Invalid JSON"}`, rr.Body.String(), "should return invalid json message")
}

func (s *TestSuperheroesControllerSuite) TestCreateInvalidPayload() {
	s.router.HandleFunc("/superheroes", s.controller.Create).Methods("POST")

	rr := s.ServeRequest("POST", "/superheroes", bytes.NewBufferString(`{ "Name": "" }`))

	s.Equal(400, rr.Code, "should return status 400 bad request")
	s.Equal(`{"detail":"Bad request. Invalid payload"}`, rr.Body.String(), "should return invalid payload message")
}

func (s *TestSuperheroesControllerSuite) TestCreate() {
	s.router.HandleFunc("/superheroes", s.controller.Create).Methods("POST")

	s.repo.On("Insert", mock.Anything).Return(nil)

	rr := s.ServeRequest("POST", "/superheroes", bytes.NewBufferString(`{ "Name": "Batman" }`))

	s.Equal(rr.Code, 201, "should return status 201 created")
	s.repo.AssertExpectations(s.T())
}

func (s *TestSuperheroesControllerSuite) TestDeleteInvalidUUID() {
	s.router.HandleFunc("/superheroes/{uuid}", s.controller.Delete).Methods("DELETE")

	rr := s.ServeRequest("DELETE", "/superheroes/1234", nil)

	s.Equal(http.StatusBadRequest, rr.Code, "should return status 400 Bad Request")
	s.Equal(`{"detail":"Bad request. Invalid UUID."}`, rr.Body.String(), "should return empty body")
}

func (s *TestSuperheroesControllerSuite) TestDelete() {
	s.router.HandleFunc("/superheroes/{uuid}", s.controller.Delete).Methods("DELETE")

	shId := uuid.NewV4()
	s.repo.On("Delete", shId).Return(nil)

	rr := s.ServeRequest("DELETE", "/superheroes/"+shId.String(), nil)

	s.Equal(http.StatusNoContent, rr.Code, "should return status 204 No Content")
	s.repo.AssertExpectations(s.T())
}

func TestControllerSuite(t *testing.T) {
	suite.Run(t, new(TestSuperheroesControllerSuite))
}
