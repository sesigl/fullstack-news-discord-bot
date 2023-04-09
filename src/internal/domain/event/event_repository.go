package event

import (
	"fmt"
	"github.com/google/uuid"
)

type NotFoundError struct {
	Id uuid.UUID
}

func (e NotFoundError) Error() string {
	return fmt.Sprintf("could not find an event with id: %v", e.Id)
}

type Repository interface {
	Save(event Event) (Event, error)
	DeleteByID(id uuid.UUID) error
	FindByID(id uuid.UUID) (Event, error)
}
