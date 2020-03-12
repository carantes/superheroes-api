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
	service    *SuperheroesServiceMock
}

// Mocks
type SuperheroRepositoryMock struct {
	mock.Mock
}

func (mock *SuperheroRepositoryMock) FindAll(filter *superheroesbundle.Superhero) ([]superheroesbundle.Superhero, error) {
	args := mock.Called(filter)
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

type SuperheroesServiceMock struct {
	mock.Mock
}

func (mock *SuperheroesServiceMock) Search(name string) (superheroesbundle.SuperheroAPISearchResponse, error) {
	args := mock.Called(name)
	return args.Get(0).(superheroesbundle.SuperheroAPISearchResponse), args.Error(1)
}

func (s *TestSuperheroesControllerSuite) SetupSuite() {
	s.repo = new(SuperheroRepositoryMock)
	s.service = new(SuperheroesServiceMock)
	s.controller = superheroesbundle.NewSuperheroesController(s.repo, s.service)
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
		*superheroesbundle.NewSuperhero("Batman", "Bruce Wayne", superheroesbundle.GoodAlignment, 100, 47, "-", "https://www.superherodb.com/pictures2/portraits/10/100/10441.jpg", 10),
		*superheroesbundle.NewSuperhero("Wolverine", "Logan", superheroesbundle.GoodAlignment, 63, 89, "Adventurer, instructor, former bartender...", "https://www.superherodb.com/pictures2/portraits/10/100/161.jpg", 7),
	}

	s.repo.On("FindAll", &superheroesbundle.Superhero{}).Return(mockData, nil)

	// Setup handler for Index function
	s.router.HandleFunc("/superheroes", s.controller.Index).Methods("GET")

	// Call index
	rr := s.ServeRequest("GET", "/superheroes", nil)

	// Assert return expected result
	s.Equal(http.StatusOK, rr.Code, "should return status 200 OK")
}

func (s *TestSuperheroesControllerSuite) TestGet() {

	s.Run("invalid UUID", func() {
		s.router.HandleFunc("/superheroes/{uuid}", s.controller.Get).Methods("GET")

		rr := s.ServeRequest("GET", "/superheroes/1234", nil)

		s.Equal(http.StatusBadRequest, rr.Code, "should return status 400 Bad Request")
		s.Equal(`{"detail":"Bad request. Invalid UUID."}`, rr.Body.String(), "should return invalid UUID message")
	})

	s.Run("success", func() {
		s.router.HandleFunc("/superheroes/{uuid}", s.controller.Get).Methods("GET")

		mockData := *superheroesbundle.NewSuperhero("Batman", "Bruce Wayne", superheroesbundle.GoodAlignment, 100, 47, "-", "https://www.superherodb.com/pictures2/portraits/10/100/10441.jpg", 10)
		s.repo.On("FindOne", mockData.ID).Return(mockData, nil)

		rr := s.ServeRequest("GET", "/superheroes/"+mockData.ID.String(), nil)

		// Assert return expected result
		s.Equal(http.StatusOK, rr.Code, "should return status 200 success")
	})
}

func (s *TestSuperheroesControllerSuite) TestCreate() {

	s.Run("empty payload", func() {
		s.router.HandleFunc("/superheroes", s.controller.Create).Methods("POST")

		rr := s.ServeRequest("POST", "/superheroes", bytes.NewBufferString(``))

		s.Equal(http.StatusBadRequest, rr.Code, "should return status 400 bad request")
		s.Equal(`{"detail":"Bad request. Invalid JSON"}`, rr.Body.String(), "should return invalid json message")
	})

	s.Run("invalid payload", func() {
		s.router.HandleFunc("/superheroes", s.controller.Create).Methods("POST")

		rr := s.ServeRequest("POST", "/superheroes", bytes.NewBufferString(`{ "Name": "" }`))

		s.Equal(http.StatusBadRequest, rr.Code, "should return status 400 bad request")
		s.Equal(`{"detail":"Bad request. Invalid payload"}`, rr.Body.String(), "should return invalid payload message")
	})

	s.Run("not found", func() {
		s.router.HandleFunc("/superheroes", s.controller.Create).Methods("POST")

		s.repo.On("FindAll", &superheroesbundle.Superhero{Name: "Mussum"}).Return([]superheroesbundle.Superhero{}, nil)
		s.service.On("Search", "Mussum").Return(superheroesbundle.SuperheroAPISearchResponse{}, nil)

		rr := s.ServeRequest("POST", "/superheroes", bytes.NewBufferString(`{ "Name": "Mussum" }`))

		s.Equal(http.StatusNotFound, rr.Code, "should return status 404 not found")
	})

	s.Run("return existent", func() {
		s.router.HandleFunc("/superheroes", s.controller.Create).Methods("POST")

		mockData := []superheroesbundle.Superhero{
			*superheroesbundle.NewSuperhero("Batman", "Bruce Wayne", superheroesbundle.GoodAlignment, 100, 47, "-", "https://www.superherodb.com/pictures2/portraits/10/100/10441.jpg", 10),
		}
		s.repo.On("FindAll", &superheroesbundle.Superhero{Name: "Batman"}).Return(mockData, nil)

		rr := s.ServeRequest("POST", "/superheroes", bytes.NewBufferString(`{ "Name": "Batman" }`))

		s.Equal(http.StatusOK, rr.Code, "should return status 200 ok")
	})

	s.Run("created", func() {
		s.router.HandleFunc("/superheroes", s.controller.Create).Methods("POST")

		// Mock db to return nothing
		s.repo.On("FindAll", &superheroesbundle.Superhero{Name: "Wolverine"}).Return([]superheroesbundle.Superhero{}, nil)

		// Mock service to return a match
		s.service.On("Search", "Wolverine").Return(superheroesbundle.SuperheroAPISearchResponse{
			[]superheroesbundle.SuperheroAPISearchResponseResult{
				superheroesbundle.SuperheroAPISearchResponseResult{Name: "Wolverine"},
			},
		}, nil)

		//Mock insert to return success
		s.repo.On("Insert", mock.Anything).Return(nil)

		rr := s.ServeRequest("POST", "/superheroes", bytes.NewBufferString(`{ "Name": "Wolverine" }`))

		s.Equal(http.StatusCreated, rr.Code, "should return status 201 created")
	})

}

func (s *TestSuperheroesControllerSuite) TestDelete() {

	s.Run("success", func() {
		s.router.HandleFunc("/superheroes/{uuid}", s.controller.Delete).Methods("DELETE")

		shId := uuid.NewV4()
		s.repo.On("Delete", shId).Return(nil)

		rr := s.ServeRequest("DELETE", "/superheroes/"+shId.String(), nil)

		s.Equal(http.StatusNoContent, rr.Code, "should return status 204 No Content")
	})

	s.Run("invalid UUID", func() {
		s.router.HandleFunc("/superheroes/{uuid}", s.controller.Delete).Methods("DELETE")

		rr := s.ServeRequest("DELETE", "/superheroes/1234", nil)

		s.Equal(http.StatusBadRequest, rr.Code, "should return status 400 Bad Request")
		s.Equal(`{"detail":"Bad request. Invalid UUID."}`, rr.Body.String(), "should return empty body")
	})
}

func (s *TestSuperheroesControllerSuite) AfterTest(_, _ string) {
	s.repo.AssertExpectations(s.T())
}

func TestControllerSuite(t *testing.T) {
	suite.Run(t, new(TestSuperheroesControllerSuite))
}
