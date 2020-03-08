package core

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

//Base contains common column for all tables
type Base struct {
	ID        uuid.UUID         `json:"id" gorm:"type:uuid;primary_key;"`
	CreatedAt time.Time         `json:"created_at"`
	UpdatedAt time.Time         `json:"updated_at"`
	DeletedAt *time.Time        `json:"-" sql:"index"`
	Errors    map[string]string `json:"-" gorm:"-"`
}
