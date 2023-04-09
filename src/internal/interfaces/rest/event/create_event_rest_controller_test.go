package event_test

import (
	"encoding/json"
	"fmt"
	"fullStack-news-discord-bot/src/internal"
	restEvent "fullStack-news-discord-bot/src/internal/interfaces/rest/event"
	"fullStack-news-discord-bot/src/test/fake/repository"
	lambdaEvent "github.com/aws/aws-lambda-go/events"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestCreateEvent_createsEvent(t *testing.T) {
	eventRepository, createEventRestController := createEventRestControllerMocksInjected()

	createEventRequestObject := createBasicCreateEventRequestObject()

	request := parameterRequestObjectToLambdaRequestObject(createEventRequestObject, "POST")

	_, _ = createEventRestController.HandleLambdaEvent(request)

	existingEvents := eventRepository.FindAll()

	assert.Len(t, existingEvents, 1)
	assert.NotEqual(t, existingEvents[0].GetID(), uuid.Nil)
	assert.Equal(t, existingEvents[0].GetTopic(), createEventRequestObject.Topic)
	assert.Equal(t, existingEvents[0].GetName(), createEventRequestObject.Name)
	assert.Equal(t, existingEvents[0].GetImageSrc(), createEventRequestObject.ImageSrc)
	assert.True(t, existingEvents[0].GetStart().Equal(createEventRequestObject.Start))
	assert.True(t, existingEvents[0].GetEnd().Equal(createEventRequestObject.End))
}

func TestCreateEvent_whenRequiredParametersAreMissing_returns4xx(t *testing.T) {
	_, createEventRestController := createEventRestControllerMocksInjected()

	createEventRequestObject := createBasicCreateEventRequestObject()
	createEventRequestObject.Name = ""

	request := parameterRequestObjectToLambdaRequestObject(createEventRequestObject, "POST")

	response, _ := createEventRestController.HandleLambdaEvent(request)

	assert.Equal(t, 400, response.StatusCode)
	assert.Contains(t, response.Body, "Field validation for 'Name' failed on the 'required'")
}

func TestDeleteEvent_deletesEvent(t *testing.T) {
	eventRepository, createEventRestController := createEventRestControllerMocksInjected()

	createEventRequestObject := createBasicCreateEventRequestObject()
	createEventRequest := parameterRequestObjectToLambdaRequestObject(createEventRequestObject, "POST")
	createdEventResponse, _ := createEventRestController.HandleLambdaEvent(createEventRequest)

	responseData := restEvent.CreateEventResponse{}
	_ = json.Unmarshal([]byte(createdEventResponse.Body), &responseData)

	deleteLambdaRequestObject := createLambdaRequestObject("DELETE", fmt.Sprintf("/%v", responseData.ID))
	_, _ = createEventRestController.HandleLambdaEvent(deleteLambdaRequestObject)

	existingEvents := eventRepository.FindAll()
	assert.Len(t, existingEvents, 0)
}

func TestDeleteEvent_returns404IfNotExists(t *testing.T) {
	_, createEventRestController := createEventRestControllerMocksInjected()

	deleteLambdaRequestObject := createLambdaRequestObject("DELETE", fmt.Sprintf("/%v", uuid.New()))
	deleteResponse, _ := createEventRestController.HandleLambdaEvent(deleteLambdaRequestObject)

	assert.Equal(t, deleteResponse.StatusCode, 404)
}

func TestUpdateEvent_notExisting_returns404(t *testing.T) {
	_, createEventRestController := createEventRestControllerMocksInjected()

	putLambdaRequestObject := createLambdaRequestObject("PUT", fmt.Sprintf("/%v", uuid.New()))
	putResponse, _ := createEventRestController.HandleLambdaEvent(putLambdaRequestObject)

	assert.Equal(t, putResponse.StatusCode, 404)
}

func TestUpdateEvent_existing_updatesAndReturns200(t *testing.T) {
	eventRepository, createEventRestController := createEventRestControllerMocksInjected()

	putLambdaRequestObject := createLambdaRequestObject("PUT", fmt.Sprintf("/%v", uuid.New()))

	createEventRequestObject := createBasicCreateEventRequestObject()
	createEventRequestObject.Name = "newName"
	createEventRequestJson, _ := json.Marshal(createEventRequestObject)
	putLambdaRequestObject.Body = string(createEventRequestJson)

	putResponse, _ := createEventRestController.HandleLambdaEvent(putLambdaRequestObject)

	existingEvents := eventRepository.FindAll()

	assert.Len(t, existingEvents, 1)
	assert.Equal(t, existingEvents[0].GetName(), createEventRequestObject.Name)
	assert.Equal(t, putResponse.StatusCode, 200)
}

func createEventRestControllerMocksInjected() (repository.FakeEventRepository, restEvent.CreateEventRestController) {
	var eventRepository = repository.NewFakeEventRepository()
	createEventRestController := internal.InitializeCreateEventRestControllerWithMocks(eventRepository)
	return eventRepository, createEventRestController
}

func createBasicCreateEventRequestObject() restEvent.CreateEventRequestObject {
	now := time.Now()
	createEventRequestObject := restEvent.CreateEventRequestObject{
		Topic:    "topic",
		Name:     "name",
		ImageSrc: "imageSrc",
		Start:    now,
		End:      now.Add(time.Hour * 1),
	}
	return createEventRequestObject
}

func parameterRequestObjectToLambdaRequestObject(
	createEventRequestObject restEvent.CreateEventRequestObject,
	requestMethod string,
) lambdaEvent.LambdaFunctionURLRequest {
	jsonBody, _ := json.Marshal(createEventRequestObject)

	request := lambdaEvent.LambdaFunctionURLRequest{
		Body: string(jsonBody),
		RequestContext: lambdaEvent.LambdaFunctionURLRequestContext{
			HTTP: lambdaEvent.LambdaFunctionURLRequestContextHTTPDescription{
				Method: requestMethod,
			},
		},
	}
	return request
}

func createLambdaRequestObject(
	requestMethod string,
	path string,
) lambdaEvent.LambdaFunctionURLRequest {
	request := lambdaEvent.LambdaFunctionURLRequest{
		RequestContext: lambdaEvent.LambdaFunctionURLRequestContext{
			HTTP: lambdaEvent.LambdaFunctionURLRequestContextHTTPDescription{
				Method: requestMethod,
				Path:   path,
			},
		},
	}
	return request
}
