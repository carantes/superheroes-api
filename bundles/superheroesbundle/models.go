package superheroesbundle

import uuid "github.com/satori/go.uuid"

// Superhero model
type Superhero struct {
	ID           uuid.UUID         `json:"id"`
	Name         string            `json:"name"`
	FullName     string            `json:"fullname"`
	Intelligence int               `json:"intelligence"`
	Power        int               `json:"power"`
	Occupation   string            `json:"occupation"`
	Errors       map[string]string `json:"-"`
}

// NewSuperhero create a new Superhero
func NewSuperhero(name string, fullName string, intelligence int, power int, occupation string) *Superhero {
	return &Superhero{
		ID:           uuid.NewV4(),
		Name:         name,
		FullName:     fullName,
		Intelligence: intelligence,
		Power:        power,
		Occupation:   occupation,
	}
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
