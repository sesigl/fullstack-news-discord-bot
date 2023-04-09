package event

import (
	"github.com/google/uuid"
	"time"
)

type Event struct {
	id       uuid.UUID
	topic    string
	name     string
	imageSrc string

	start time.Time
	end   time.Time
}

func (event *Event) GetID() uuid.UUID {
	return event.id
}

func (event *Event) GetTopic() string {
	return event.topic
}

func (event *Event) GetName() string {
	return event.name
}

func (event *Event) GetStart() time.Time {
	return event.start
}

func (event *Event) GetEnd() time.Time {
	return event.end
}

func (event *Event) GetImageSrc() string {
	return event.imageSrc
}

func (event *Event) SetID(id uuid.UUID) {
	event.id = id
}
