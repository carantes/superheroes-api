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
	newRows := sqlmock.NewRows([]string{"id", "name", "full_name", "alignment", "intelligence", "power", "occupation", "image_url", "relatives", "created_at", "updated_at", "deleted_at"})

	for _, sh := range heroes {
		newRows.AddRow(sh.ID.String(), sh.Name, sh.FullName, sh.Alignment, sh.Intelligence, sh.Power, sh.Occupation, sh.ImageURL, sh.TotalRelatives, sh.CreatedAt, sh.UpdatedAt, sh.DeletedAt)
	}

	return newRows
}

func (s *TestSuperheroesRepositorySuite) CreateSuperheroGroupDBRows(groups []superheroesbundle.SuperheroGroup) *sqlmock.Rows {
	newRows := sqlmock.NewRows([]string{"id", "name", "created_at", "updated_at", "deleted_at"})

	for _, sh := range groups {
		newRows.AddRow(sh.ID.String(), sh.Name, sh.CreatedAt, sh.UpdatedAt, sh.DeletedAt)
	}

	return newRows
}

func (s *TestSuperheroesRepositorySuite) TestFindAll() {

	s.Run("return all", func() {
		mockData := []superheroesbundle.Superhero{
			*superheroesbundle.NewSuperhero("Batman", "Bruce Wayne", superheroesbundle.GoodAlignment, 100, 47, "-", "http://www.hero.com/161.jpg", 10),
			*superheroesbundle.NewSuperhero("Wolverine", "Logan", superheroesbundle.GoodAlignment, 63, 89, "Adventurer, instructor...", "http://www.hero.com/161.jpg", 7),
		}

		s.mock.ExpectQuery(regexp.QuoteMeta(
			`SELECT * FROM "superheroes" WHERE "superheroes"."deleted_at" IS NULL`)).
			WillReturnRows(s.CreateSuperheroDBRows(mockData))

		sh, err := s.repo.FindAll(&superheroesbundle.Superhero{})
		s.NoError(err, "should not return error on find all")
		s.Len(sh, 2)
		s.Equal(mockData, sh)
	})
}

func (s *TestSuperheroesRepositorySuite) TestFindOne() {
	shMock := *superheroesbundle.NewSuperhero("Batman", "Bruce Wayne", superheroesbundle.GoodAlignment, 100, 47, "-", "http://www.hero.com/161.jpg", 10)
	shGroupMock := *shMock.AddGroup("Justice League")

	s.mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT * FROM "superheroes" WHERE "superheroes"."deleted_at" IS NULL
			AND ((id = $1)) ORDER BY "superheroes"."id" ASC LIMIT 1`)).
		WithArgs(shMock.ID).
		WillReturnRows(s.CreateSuperheroDBRows([]superheroesbundle.Superhero{shMock}))
	s.mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT * FROM "superhero_groups"  WHERE "superhero_groups"."deleted_at" IS NULL
			AND (("superhero_id" IN ($1))) ORDER BY "superhero_groups"."id" ASC`)).
		WithArgs(shMock.ID).
		WillReturnRows(s.CreateSuperheroGroupDBRows([]superheroesbundle.SuperheroGroup{shGroupMock}))

	sh, err := s.repo.FindOne(shMock.ID)
	s.NoError(err, "should not return error on Find one")
	s.Equal(shMock, sh)
}

func (s *TestSuperheroesRepositorySuite) TestCreate() {
	mockData := superheroesbundle.NewSuperhero("Wolverine", "Logan", superheroesbundle.GoodAlignment, 63, 89, "Adventurer, instructor...", "http://www.hero.com/161.jpg", 7)

	s.mock.ExpectBegin()
	s.mock.ExpectQuery(regexp.QuoteMeta(
		`INSERT INTO "superheroes" ("id","created_at","updated_at","deleted_at","name","full_name","alignment","intelligence","power","occupation","image_url","relatives")
			VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12) RETURNING "superheroes"."id"`)).
		WithArgs(mockData.ID, AnyTime{}, AnyTime{}, nil, mockData.Name, mockData.FullName, mockData.Alignment, mockData.Intelligence, mockData.Power, mockData.Occupation, mockData.ImageURL, mockData.TotalRelatives).
		WillReturnRows(
			sqlmock.NewRows([]string{"id"}).AddRow(mockData.ID.String()))
	s.mock.ExpectCommit()

	err := s.repo.Insert(mockData)
	s.NoError(err, "should not return error on create")
}

func (s *TestSuperheroesRepositorySuite) TestDelete() {
	mockData := superheroesbundle.NewSuperhero("Wolverine", "Logan", superheroesbundle.GoodAlignment, 63, 89, "Adventurer, instructor...", "http://www.hero.com/161.jpg", 7)

	//TODO: Test 404 if not found the superhero
	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta(
		`UPDATE "superheroes" SET "deleted_at"=$1  WHERE "superheroes"."deleted_at" IS NULL AND ((id = $2))`)).
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
