package core

import (
	"fmt"

	uuid "github.com/satori/go.uuid"
)

// HTTPError is a custom error for client responses
type HTTPError struct {
	Cause  error  `json:"-"`
	Detail string `json:"detail"`
	Status int    `json:"-"`
}

func (e *HTTPError) Error() string {
	if e.Cause == nil {
		return e.Detail
	}
	return e.Detail + " : " + e.Cause.Error()
}

// NewHTTPError instance
func NewHTTPError(err error, status int, detail string) error {
	return &HTTPError{
		Cause:  err,
		Detail: detail,
		Status: status,
	}
}

// RepositoryNotFoundError is a custom error for not found on repository
type RepositoryNotFoundError struct {
	error
	ID uuid.UUID
}

func (e *RepositoryNotFoundError) Error() string {
	return fmt.Sprintf("object with UUD %s not found on repository", e.ID.String())
}

// NewRepositoryNotFoundError instance
func NewRepositoryNotFoundError(id uuid.UUID, err error) error {
	return &RepositoryNotFoundError{
		ID:    id,
		error: err,
	}
}
