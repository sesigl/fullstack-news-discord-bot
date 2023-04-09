package repository

import (
	"fullStack-news-discord-bot/src/internal/domain/event"
	"github.com/google/uuid"
	"golang.org/x/exp/maps"
)

func NewFakeEventRepository() FakeEventRepository {
	return FakeEventRepository{
		events: make(map[string]event.Event),
	}
}

type FakeEventRepository struct {
	events map[string]event.Event
}

func (repository FakeEventRepository) Save(event event.Event) (event.Event, error) {
	if event.GetID() == uuid.Nil {
		event.SetID(uuid.New())
	}

	repository.events[event.GetID().String()] = event
	return event, nil
}

func (repository FakeEventRepository) FindAll() []event.Event {
	return maps.Values(repository.events)
}

func (repository FakeEventRepository) DeleteByID(id uuid.UUID) error {
	idMapKey := id.String()

	_, exist := repository.events[idMapKey]
	if !exist {
		return event.NotFoundError{
			Id: id,
		}
	}

	delete(repository.events, idMapKey)
	return nil
}

func (repository FakeEventRepository) FindByID(id uuid.UUID) (event.Event, error) {
	idMapKey := id.String()

	existingEvent, exist := repository.events[idMapKey]
	if !exist {
		return event.Event{}, event.NotFoundError{
			Id: id,
		}
	}

	return existingEvent, nil
}
