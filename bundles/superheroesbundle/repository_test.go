package superheroesbundle_test

import (
	"database/sql"
	"database/sql/driver"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/carantes/superheroes-api/bundles/superheroesbundle"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/stretchr/testify/suite"
)

type TestSuperheroesRepositorySuite struct {
	suite.Suite
	DB   *gorm.DB
	mock sqlmock.Sqlmock
	repo *superheroesbundle.SuperheroesRepository
}

type AnyTime struct{}

// Match satisfies sqlmock.Argument interface
func (a AnyTime) Match(v driver.Value) bool {
	_, ok := v.(time.Time)
	return ok
}

func (s *TestSuperheroesRepositorySuite) SetupSuite() {
	var (
		db  *sql.DB
		err error
	)

	db, s.mock, err = sqlmock.New()
	s.NoError(err)

	s.DB, err = gorm.Open("postgres", db)
	s.NoError(err)

	s.DB.LogMode(true)

	s.repo = superheroesbundle.NewSuperheroesRepository(s.DB)
}

func (s *TestSuperheroesRepositorySuite) CreateSuperheroDBRows(heroes []superheroesbundle.Superhero) *sqlmock.Rows {
	newRows := sqlmock.NewRows([]string{"id", "name", "full_name", "intelligence", "power", "occupation", "created_at", "updated_at", "deleted_at"})

	for _, sh := range heroes {
		newRows.AddRow(sh.ID.String(), sh.Name, sh.FullName, sh.Intelligence, sh.Power, sh.Occupation, sh.CreatedAt, sh.UpdatedAt, sh.DeletedAt)
	}

	return newRows
}

func (s *TestSuperheroesRepositorySuite) TestFindAll() {
	mockData := []superheroesbundle.Superhero{
		*superheroesbundle.NewSuperhero("Batman", "Bruce Wayne", 100, 47, "-"),
		*superheroesbundle.NewSuperhero("Wolverine", "Logan", 63, 89, "Adventurer, instructor, former bartender..."),
	}

	s.mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT * FROM "superheros" WHERE "superheros"."deleted_at" IS NULL`)).
		WillReturnRows(s.CreateSuperheroDBRows(mockData))

	sh, err := s.repo.FindAll()
	s.NoError(err, "should not return error on find all")
	s.Len(sh, 2)
	s.Equal(mockData, sh)
}

func (s *TestSuperheroesRepositorySuite) TestFindOne() {

	//TODO: Test 404 if not found the superhero
	mockData := []superheroesbundle.Superhero{
		*superheroesbundle.NewSuperhero("Batman", "Bruce Wayne", 100, 47, "-"),
	}

	s.mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT * FROM "superheros" WHERE "superheros"."deleted_at" IS NULL
			AND ((id = $1)) ORDER BY "superheros"."id" ASC LIMIT 1`)).
		WithArgs(mockData[0].ID).
		WillReturnRows(s.CreateSuperheroDBRows(mockData))

	sh, err := s.repo.FindOne(mockData[0].ID)
	s.NoError(err, "should not return error on Find one")
	s.Equal(mockData[0], sh)
}

func (s *TestSuperheroesRepositorySuite) TestCreate() {
	mockData := superheroesbundle.NewSuperhero("Wolverine", "Logan", 63, 89, "Adventurer, instructor, former bartender...")

	s.mock.ExpectBegin()
	s.mock.ExpectQuery(regexp.QuoteMeta(
		`INSERT INTO "superheros" ("id","created_at","updated_at","deleted_at","name","full_name","intelligence","power","occupation")
			VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9) RETURNING "superheros"."id"`)).
		WithArgs(mockData.ID, AnyTime{}, AnyTime{}, nil, mockData.Name, mockData.FullName, mockData.Intelligence, mockData.Power, mockData.Occupation).
		WillReturnRows(
			sqlmock.NewRows([]string{"id"}).AddRow(mockData.ID.String()))
	s.mock.ExpectCommit()

	err := s.repo.Insert(mockData)
	s.NoError(err, "should not return error on create")
}

func (s *TestSuperheroesRepositorySuite) TestDelete() {
	mockData := superheroesbundle.NewSuperhero("Wolverine", "Logan", 63, 89, "Adventurer, instructor, former bartender...")

	//TODO: Test 404 if not found the superhero
	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta(
		`UPDATE "superheros" SET "deleted_at"=$1  WHERE "superheros"."deleted_at" IS NULL AND ((id = $2))`)).
		WithArgs(AnyTime{}, mockData.ID).
		WillReturnResult(sqlmock.NewResult(0, 1))
	s.mock.ExpectCommit()

	err := s.repo.Delete(mockData.ID)
	s.NoError(err, "should not return error on delete")
}

func (s *TestSuperheroesRepositorySuite) AfterTest(_, _ string) {
	s.NoError(s.mock.ExpectationsWereMet())
}

func TestRepositorySuite(t *testing.T) {
	suite.Run(t, new(TestSuperheroesRepositorySuite))
}
