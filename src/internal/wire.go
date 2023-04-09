//go:build wireinject
// +build wireinject

package internal

import (
	"fullStack-news-discord-bot/src/internal/application"
	"fullStack-news-discord-bot/src/internal/domain/event"
	"fullStack-news-discord-bot/src/internal/infrastructure/discord"
	restEvent "fullStack-news-discord-bot/src/internal/interfaces/rest/event"
	"github.com/google/wire"
)

var providerSet wire.ProviderSet = wire.NewSet(
	application.NewEventApplicationService,
	discord.NewDiscordEventRepository,
	restEvent.NewCreateEventRestController,

	wire.Bind(new(event.Repository), new(discord.EventRepository)),
)

func InitializeEventRepository() event.Repository {
	wire.Build(providerSet)
	return discord.EventRepository{}
}

func InitializeEventApplicationService() application.EventApplicationService {
	wire.Build(providerSet)
	return application.EventApplicationService{}
}

func InitializeCreateEventRestController() restEvent.CreateEventRestController {
	wire.Build(providerSet)
	return restEvent.CreateEventRestController{}
}

var providerSetForTesting wire.ProviderSet = wire.NewSet(
	application.NewEventApplicationService,
	discord.NewDiscordEventRepository,
	restEvent.NewCreateEventRestController,
)

func InitializeCreateEventRestControllerWithMocks(mockEventRepository event.Repository) restEvent.CreateEventRestController {
	wire.Build(providerSetForTesting)
	return restEvent.CreateEventRestController{}
}
