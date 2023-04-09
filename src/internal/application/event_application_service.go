package application

import (
	"fullStack-news-discord-bot/src/internal/domain/event"
	"github.com/google/uuid"
	"time"
)

type CreateEventCommand struct {
	Topic    string
	Name     string
	ImageSrc string

	Start time.Time
	End   time.Time
}

type UpdateEventCommand struct {
	Id       uuid.UUID
	Topic    string
	Name     string
	ImageSrc string

	Start time.Time
	End   time.Time
}

func NewEventApplicationService(eventRepository event.Repository) EventApplicationService {
	return EventApplicationService{
		eventRepository: eventRepository,
	}
}

type EventApplicationService struct {
	eventRepository event.Repository
}

func (as EventApplicationService) CreateEventUseCase(command CreateEventCommand) (event.Event, error) {
	createdEvent, err := event.CreateEvent(event.CreateParams{
		Topic:    command.Topic,
		Name:     command.Name,
		ImageSrc: command.ImageSrc,
		Start:    command.Start,
		End:      command.End,
	})

	if err != nil {
		return event.Event{}, err
	}

	persistedEvent, err := as.eventRepository.Save(createdEvent)

	return persistedEvent, err
}

func (as EventApplicationService) DeleteEventUseCase(id uuid.UUID) error {
	return as.eventRepository.DeleteByID(id)
}

func (as EventApplicationService) UpdateEventUseCase(command UpdateEventCommand) (event.Event, error) {
	_, findByIdErr := as.eventRepository.FindByID(command.Id)
	if findByIdErr != nil {
		return event.Event{}, findByIdErr
	}

	updatedEvent, createEventErr := event.CreateEvent(event.CreateParams{
		Id:       command.Id.String(),
		Topic:    command.Topic,
		Name:     command.Name,
		ImageSrc: command.ImageSrc,
		Start:    command.Start,
		End:      command.End,
	})
	if createEventErr != nil {
		return event.Event{}, createEventErr
	}

	persistedEvent, saveErr := as.eventRepository.Save(updatedEvent)
	if saveErr != nil {
		return event.Event{}, saveErr
	}

	return persistedEvent, findByIdErr
}
