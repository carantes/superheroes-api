package superheroesbundle

import (
	"errors"
	"strings"

	"github.com/carantes/superheroes-api/core"
	uuid "github.com/satori/go.uuid"
)

// SuperheroAlignment is the list of options for alignment
type SuperheroAlignment int

// SuperheroAlignment list (Hero/Vilain)
const (
	UndefinedAlignment SuperheroAlignment = iota
	GoodAlignment
	BadAlignment
)

// ToString return a string from a SuperheroAlignment enum
func (a SuperheroAlignment) ToString() string {
	names := []string{
		"",
		"good",
		"bad",
	}
	return names[a]
}

// FromString return a SuperheroAlignment from a string
func (a *SuperheroAlignment) FromString(s string) SuperheroAlignment {
	switch strings.ToUpper(s) {
	case "GOOD":
		return GoodAlignment
	case "BAD":
		return BadAlignment
	}

	return UndefinedAlignment
}

// SuperheroGroup association
type SuperheroGroup struct {
	core.Base
	Name        string `json:"name" gorm:"size:128;not null;"`
	SuperheroID uuid.UUID
}

// Superhero model
type Superhero struct {
	core.Base
	Name           string             `json:"name" gorm:"size:128;not null;"`
	FullName       string             `json:"fullname" gorm:"size:256;"`
	Alignment      SuperheroAlignment `json:"alignment"`
	Intelligence   int                `json:"intelligence"`
	Power          int                `json:"power"`
	Occupation     string             `json:"occupation" gorm:"size:128;"`
	ImageURL       string             `json:"imageurl" gorm:"size:256;"`
	TotalRelatives int                `json:"relatives" gorm:"column:relatives;"`
	Groups         []SuperheroGroup   `json:"groups,omitempty"`
}

// TableName changed to be `superheroes`
func (Superhero) TableName() string {
	return "superheroes"
}

// NewSuperhero create a new Superhero
func NewSuperhero(name string, fullName string, alignment SuperheroAlignment, intelligence int, power int, occupation string, imageURL string, relatives int) *Superhero {
	sh := &Superhero{
		Name:           name,
		FullName:       fullName,
		Alignment:      alignment,
		Intelligence:   intelligence,
		Power:          power,
		Occupation:     occupation,
		ImageURL:       imageURL,
		TotalRelatives: relatives,
	}

	sh.ID = uuid.NewV4()
	return sh
}

// AddGroup append new group to superhero groups
func (sh *Superhero) AddGroup(name string) *SuperheroGroup {
	group := SuperheroGroup{Name: name}
	group.ID = uuid.NewV4()
	sh.Groups = append(sh.Groups, group)
	return &group
}

// Validate Superhero
func (sh *Superhero) Validate() error {

	if sh.Name == "" {
		return errors.New("superhero must have a name")
	}

	return nil
}
