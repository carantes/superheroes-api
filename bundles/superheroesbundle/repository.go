package superheroesbundle

import (
	"fmt"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

// SuperheroesRepository handle db functions
type SuperheroesRepository struct {
	db *gorm.DB
}

// SuperheroNotFoundError is a custom error for not found on repository
type SuperheroNotFoundError struct {
	error
	ID uuid.UUID
}

func (e *SuperheroNotFoundError) Error() string {
	return fmt.Sprintf("superhero with UUD %s not found", e.ID.String())
}

// NewSuperheroesRepository instance
func NewSuperheroesRepository(db *gorm.DB) *SuperheroesRepository {
	return &SuperheroesRepository{
		db: db,
	}
}

// FindAll implement SuperheroesRepositoryInterface
func (repo *SuperheroesRepository) FindAll() ([]Superhero, error) {
	var shs []Superhero

	repo.db.Find(&shs)

	return shs, nil
}

// FindOne implement SuperheroesRepositoryInterface
func (repo *SuperheroesRepository) FindOne(id uuid.UUID) (Superhero, error) {
	var sh Superhero

	err := repo.db.First(&sh, "id = ?", id).Error

	if err != nil {
		return sh, &SuperheroNotFoundError{ID: id, error: err}
	}

	return sh, nil
}

// Insert implement SuperheroesRepositoryInterface
func (repo *SuperheroesRepository) Insert(sh *Superhero) error {
	return repo.db.Create(sh).Error
}

// Delete implement SuperheroesRepositoryInterface
func (repo *SuperheroesRepository) Delete(id uuid.UUID) error {

	err := repo.db.Delete(&Superhero{}, "id = ?", id).Error

	if err != nil {
		return &SuperheroNotFoundError{ID: id, error: err}
	}

	return nil
}
