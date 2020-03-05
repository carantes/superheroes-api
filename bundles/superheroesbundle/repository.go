package superheroesbundle

import (
	"fmt"

	uuid "github.com/satori/go.uuid"
)

// SuperheroesRepository handle db functions
type SuperheroesRepository struct {
	superheroes []Superhero
}

// SuperheroNotFoundError is a custom error for not found on repository
type SuperheroNotFoundError struct {
	ID uuid.UUID
}

func (e *SuperheroNotFoundError) Error() string {
	return fmt.Sprintf("superhero with UUD %s not found", e.ID.String())
}

// NewSuperheroesRepository instance
func NewSuperheroesRepository(data []Superhero) *SuperheroesRepository {
	return &SuperheroesRepository{
		superheroes: data,
	}
}

// FindAll implement SuperheroesRepositoryInterface
func (repo *SuperheroesRepository) FindAll() ([]Superhero, error) {
	return repo.superheroes, nil
}

// FindOne implement SuperheroesRepositoryInterface
func (repo *SuperheroesRepository) FindOne(id uuid.UUID) (Superhero, error) {
	for _, sh := range repo.superheroes {
		if sh.ID == id {
			return sh, nil
		}
	}

	return Superhero{}, &SuperheroNotFoundError{ID: id}
}

// Insert implement SuperheroesRepositoryInterface
func (repo *SuperheroesRepository) Insert(sh *Superhero) error {
	sh.ID = uuid.NewV4()
	repo.superheroes = append(repo.superheroes, *sh)
	return nil
}

// Delete implement SuperheroesRepositoryInterface
func (repo *SuperheroesRepository) Delete(id uuid.UUID) error {
	for i, sh := range repo.superheroes {
		if sh.ID == id {
			repo.superheroes = append(repo.superheroes[:i], repo.superheroes[i+1:]...)
			return nil
		}
	}

	return &SuperheroNotFoundError{ID: id}
}
