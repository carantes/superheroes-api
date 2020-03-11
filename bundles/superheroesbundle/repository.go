package superheroesbundle

import (
	"strings"

	"github.com/carantes/superheroes-api/core"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

// SuperheroesRepository handle db functions
type SuperheroesRepository struct {
	db *gorm.DB
}

// NewSuperheroesRepository instance
func NewSuperheroesRepository(db *gorm.DB) *SuperheroesRepository {
	return &SuperheroesRepository{
		db: db,
	}
}

// FindAll implement SuperheroesRepositoryInterface
func (repo *SuperheroesRepository) FindAll(filter *Superhero) ([]Superhero, error) {
	var shs []Superhero

	// Query builder
	tx := repo.db.Model(&Superhero{})

	if filter.Name != "" {
		tx = tx.Where("UPPER(name) LIKE ?", "%"+strings.ToUpper(filter.Name)+"%")
	}

	if filter.Alignment != UndefinedAlignment {
		tx = tx.Where("alignment = ?", filter.Alignment)
	}

	tx.Find(&shs)

	return shs, nil
}

// FindOne implement SuperheroesRepositoryInterface
func (repo *SuperheroesRepository) FindOne(id uuid.UUID) (Superhero, error) {
	var sh Superhero

	// Find by ID and preload groups list
	err := repo.db.Preload("Groups").First(&sh, "id = ?", id).Error

	if err != nil {
		return sh, core.NewRepositoryNotFoundError(id, err)
	}

	return sh, nil
}

// Insert implement SuperheroesRepositoryInterface
func (repo *SuperheroesRepository) Insert(sh *Superhero) error {
	return repo.db.Create(sh).Error
}

// Delete implement SuperheroesRepositoryInterface
func (repo *SuperheroesRepository) Delete(id uuid.UUID) error {

	result := repo.db.Delete(&Superhero{}, "id = ?", id)

	if result.Error != nil || result.RowsAffected == 0 {
		return core.NewRepositoryNotFoundError(id, result.Error)
	}

	return nil
}
