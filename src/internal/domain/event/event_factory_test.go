package event_test

import (
	"fullStack-news-discord-bot/src/internal/domain/event"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestEventFactory_CreateEvent_mapsFields(t *testing.T) {
	createParams := newEventCreateParams()
	createdEvent, err := event.CreateEvent(createParams)

	if err != nil {
		t.Errorf("Expected event topic to be nil, got %s", err)
	}

	assert.NoError(t, err)
	assert.Equal(t, createParams.Topic, createdEvent.GetTopic())
	assert.Equal(t, createParams.Name, createdEvent.GetName())
	assert.Equal(t, createParams.Start, createdEvent.GetStart())
	assert.Equal(t, createParams.End, createdEvent.GetEnd())
	assert.Equal(t, createParams.ImageSrc, createdEvent.GetImageSrc())
	assert.Zero(t, createdEvent.GetID())
}

func newEventCreateParams() event.CreateParams {
	return event.CreateParams{
		Topic:    "topic",
		Name:     "name",
		ImageSrc: "imgSrc",
		Start:    time.Now(),
		End:      time.Now().Add(time.Minute),
	}
}

func TestCreateEvent_mustContainTopic(t *testing.T) {
	createParams := newEventCreateParams()

	createParams.Topic = ""
	_, err := event.CreateEvent(createParams)

	assert.ErrorContains(t, err, "topic must be defined")
}

func TestCreateEvent_mustContainName(t *testing.T) {
	createParams := newEventCreateParams()

	createParams.Name = ""
	_, err := event.CreateEvent(createParams)

	assert.ErrorContains(t, err, "name must be defined")
}

func TestCreateEvent_mustContainImgSrc(t *testing.T) {
	createParams := newEventCreateParams()

	createParams.ImageSrc = ""
	_, err := event.CreateEvent(createParams)

	assert.ErrorContains(t, err, "imageSrc must be defined")
}

func TestCreateEvent_mustContainNonZeroStartDatetime(t *testing.T) {
	createParams := newEventCreateParams()

	createParams.Start = time.Time{}
	_, err := event.CreateEvent(createParams)

	assert.ErrorContains(t, err, "start must not be zero")
}

func TestCreateEvent_mustContainNonZeroEndDatetime(t *testing.T) {
	createParams := newEventCreateParams()

	createParams.End = time.Time{}
	_, err := event.CreateEvent(createParams)

	assert.ErrorContains(t, err, "end must not be zero")
}

func TestCreateEvent_startMustBeBeforeEnd(t *testing.T) {
	createParams := newEventCreateParams()

	now := time.Now()

	createParams.Start = now.Add(time.Hour * 1)
	createParams.End = now

	_, err := event.CreateEvent(createParams)

	assert.ErrorContains(t, err, "must be after end")
}
