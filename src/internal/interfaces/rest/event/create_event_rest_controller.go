package event

import (
	"encoding/json"
	"errors"
	"fmt"
	"fullStack-news-discord-bot/src/internal/application"
	"fullStack-news-discord-bot/src/internal/domain/event"
	"github.com/aws/aws-lambda-go/events"
	"github.com/go-playground/validator"
	"github.com/google/uuid"
	"strings"
	"time"
)

func NewCreateEventRestController(
	eventApplicationService application.EventApplicationService,
) CreateEventRestController {
	return CreateEventRestController{
		eventApplicationService: eventApplicationService,
	}
}

type CreateEventRestController struct {
	eventApplicationService application.EventApplicationService
}

type CreateEventRequestObject struct {
	Topic    string `json:"topic"`
	Name     string `json:"name" validate:"required,min=3,max=100"`
	ImageSrc string `json:"imageSrc" validate:"required,min=3,max=100"`

	Start time.Time `json:"start" validate:"required"`
	End   time.Time `json:"end" validate:"required"`
}

type CreateEventResponse struct {
	ID       string `json:"id"`
	Topic    string `json:"topic"`
	Name     string `json:"name"`
	ImageSrc string `json:"imageSrc"`

	Start time.Time `json:"start"`
	End   time.Time `json:"end"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

const SuccessStatus = 200
const NotFoundStatus = 404
const ClientErrorStatus = 400
const MaxUrlSplitLength = 2

func (createEventRestController CreateEventRestController) HandleLambdaEvent(
	req events.LambdaFunctionURLRequest,
) (*events.LambdaFunctionURLResponse, error) {
	method := req.RequestContext.HTTP.Method
	path := req.RequestContext.HTTP.Path

	// fmt.Printf("Request: %v", req) //nolint:forbidigo // debug

	switch method {
	case "GET":

	case "POST":
		return createEventRestController.postHandler(req)
	case "DELETE":

		uuidFromUrl, getIdFromPathErr := createEventRestController.getIdFromPath(path)
		if getIdFromPathErr != nil {
			return nil, getIdFromPathErr
		}

		deleteEventUseCaseErr := createEventRestController.eventApplicationService.DeleteEventUseCase(
			uuidFromUrl,
		)

		if deleteEventUseCaseErr != nil {
			if errors.As(deleteEventUseCaseErr, &event.NotFoundError{}) {
				return APIResponse(NotFoundStatus, nil)
			} else {
				return nil, deleteEventUseCaseErr
			}
		}

		return APIResponse(SuccessStatus, nil)

	case "PATCH":
	case "PUT":
		uuidFromUrl, getIdFromPathErr := createEventRestController.getIdFromPath(path)
		if getIdFromPathErr != nil {
			return nil, getIdFromPathErr
		}

		body := CreateEventRequestObject{}
		err := json.Unmarshal([]byte(req.Body), &body)
		if err != nil {
			return nil, err
		}

		validate := validator.New()
		if validationError := validate.Struct(body); validationError != nil {
			return APIResponse(ClientErrorStatus, ErrorResponse{Error: validationError.Error()})
		}

		updatedEvent, updateEventUseCaseErr := createEventRestController.eventApplicationService.UpdateEventUseCase(
			application.UpdateEventCommand{
				Id:       uuidFromUrl,
				Topic:    body.Topic,
				Name:     body.Name,
				ImageSrc: body.ImageSrc,
				Start:    body.Start,
				End:      body.End,
			},
		)

		if updateEventUseCaseErr != nil {
			if errors.As(updateEventUseCaseErr, &event.NotFoundError{}) {
				return APIResponse(NotFoundStatus, nil)
			} else {
				return nil, updateEventUseCaseErr
			}
		}

		response := CreateEventResponse{
			ID:       updatedEvent.GetID().String(),
			Topic:    updatedEvent.GetTopic(),
			Name:     updatedEvent.GetName(),
			ImageSrc: updatedEvent.GetImageSrc(),
			Start:    updatedEvent.GetStart(),
			End:      updatedEvent.GetEnd(),
		}

		return APIResponse(SuccessStatus, response)

	}

	return nil, errors.New("should never happen")
}

func (createEventRestController CreateEventRestController) getIdFromPath(path string) (uuid.UUID, error) {
	pathSplit := strings.Split(path, "/")
	if len(pathSplit) != MaxUrlSplitLength {
		return uuid.UUID{}, fmt.Errorf("invalid url path for GET: %v", path)
	}
	uuidFromUrl := uuid.MustParse(pathSplit[1])
	return uuidFromUrl, nil
}

func (createEventRestController CreateEventRestController) postHandler(req events.LambdaFunctionURLRequest) (*events.LambdaFunctionURLResponse, error) {

	body := CreateEventRequestObject{}
	err := json.Unmarshal([]byte(req.Body), &body)
	if err != nil {
		return nil, err
	}

	validate := validator.New()
	if validationError := validate.Struct(body); validationError != nil {
		return APIResponse(ClientErrorStatus, ErrorResponse{Error: validationError.Error()})
	}

	createdEvent, err := createEventRestController.eventApplicationService.CreateEventUseCase(
		application.CreateEventCommand{
			Topic:    body.Topic,
			Name:     body.Name,
			ImageSrc: body.ImageSrc,
			Start:    body.Start,
			End:      body.End,
		},
	)
	if err != nil {
		return nil, err
	}

	response := CreateEventResponse{
		ID:       createdEvent.GetID().String(),
		Topic:    createdEvent.GetTopic(),
		Name:     createdEvent.GetName(),
		ImageSrc: createdEvent.GetImageSrc(),
		Start:    createdEvent.GetStart(),
		End:      createdEvent.GetEnd(),
	}

	return APIResponse(SuccessStatus, response)
}

func APIResponse(status int, body interface{}) (*events.LambdaFunctionURLResponse, error) {
	resp := events.LambdaFunctionURLResponse{Headers: map[string]string{"Content-Type": "application/json"}}
	resp.StatusCode = status
	stringBody, _ := json.Marshal(body)
	resp.Body = string(stringBody)
	return &resp, nil
}
