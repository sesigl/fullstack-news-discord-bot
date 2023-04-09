package discord

import (
	"fullStack-news-discord-bot/src/internal/domain/event"
	"github.com/google/uuid"
)

func NewDiscordEventRepository() EventRepository {
	return EventRepository{}
}

type EventRepository struct {
}

func (e EventRepository) Save(event event.Event) (event.Event, error) {
	return event, nil
}

func (e EventRepository) DeleteByID(id uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}

func (e EventRepository) FindByID(id uuid.UUID) (event.Event, error) {
	//TODO implement me
	panic("implement me")
}
