package event

import (
	"errors"
	"fmt"
	"time"
)

type CreateParams struct {
	Id       string
	Topic    string
	Name     string
	ImageSrc string

	Start time.Time
	End   time.Time
}

func CreateEvent(createParams CreateParams) (Event, error) {
	event, err := validate(createParams)

	if err != nil {
		return event, err
	}

	return Event{
		topic:    createParams.Topic,
		name:     createParams.Name,
		imageSrc: createParams.ImageSrc,
		start:    createParams.Start,
		end:      createParams.End,
	}, nil
}

func validate(createParams CreateParams) (Event, error) {
	if createParams.Topic == "" {
		return Event{}, errors.New("topic must be defined")
	}

	if createParams.Name == "" {
		return Event{}, errors.New("name must be defined")
	}

	if createParams.ImageSrc == "" {
		return Event{}, errors.New("imageSrc must be defined")
	}

	if createParams.Start.IsZero() {
		return Event{}, errors.New("start must not be zero")
	}

	if createParams.End.IsZero() {
		return Event{}, errors.New("end must not be zero")
	}

	if createParams.End.Sub(createParams.Start).Milliseconds() <= 0 {
		return Event{}, fmt.Errorf("start (%v) must be after end (%v)", createParams.Start, createParams.End)
	}
	return Event{}, nil
}
