package superheroesbundle

import (
	"github.com/carantes/superheroes-api/core"
	uuid "github.com/satori/go.uuid"
)

// Superhero model
type Superhero struct {
	core.Base
	Name         string `json:"name" gorm:"column:name;size:128;not null;"`
	FullName     string `json:"fullname" gorm:"column:full_name;size:256;"`
	Intelligence int    `json:"intelligence" gorm:"column:intelligence;"`
	Power        int    `json:"power" gorm:"column:power;"`
	Occupation   string `json:"occupation" gorm:"column:occupation;size:128;"`
}

// NewSuperhero create a new Superhero
func NewSuperhero(name string, fullName string, intelligence int, power int, occupation string) *Superhero {
	sh := &Superhero{
		Name:         name,
		FullName:     fullName,
		Intelligence: intelligence,
		Power:        power,
		Occupation:   occupation,
	}

	sh.ID = uuid.NewV4()
	return sh
}

// Validate Superhero
func (sh *Superhero) Validate() bool {
	sh.Errors = make(map[string]string)

	if sh.Name == "" {
		sh.Errors["name"] = "superhero must have a name"
	}

	if len(sh.Errors) > 0 {
		return false
	}

	return true
}
