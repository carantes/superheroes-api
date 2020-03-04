package superheroesbundle

import uuid "github.com/satori/go.uuid"

// Superhero model
type Superhero struct {
	ID           uuid.UUID
	Name         string
	FullName     string
	Intelligence int
	Power        int
	Occupation   string
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
