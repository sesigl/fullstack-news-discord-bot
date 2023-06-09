// AUTOGENERATED FILE

package discord

import (
	"fmt"
	"log"
	"time"

	"github.com/bwmarrin/discordgo"
)

var (
	Token     string = ""
	BotPrefix string = "!"
)

type configStruct struct {
	Token     string `json : "Token"`
	BotPrefix string `json : "BotPrefix"`
}

func ReadConfig() error {

	return nil

}

var BotId string
var goBot *discordgo.Session

func Start() {
	goBot, err := discordgo.New("Bot " + Token)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	u, err := goBot.User("@me")

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	BotId = u.ID

	handler(goBot)

	//err = goBot.Open()
	//
	//if err != nil {
	//	fmt.Println(err.Error())
	//	return
	//}
	goBot.Close()
	fmt.Println("Bot was running !")
}

func handler(s *discordgo.Session) {

	if true {
		_, _ = s.ChannelMessageSend("1060243787412156528", "pong")
	}

	if true {
		// Define the starting time (must be in future)
		startingTime := time.Now().Add(1 * time.Hour)
		// Define the ending time (must be after starting time)
		endingTime := startingTime.Add(30 * time.Minute)
		// Save the event
		scheduledEvent, err := s.GuildScheduledEventCreate("1060243786934009967", &discordgo.GuildScheduledEventParams{
			Name:               "Amazing Event",
			Description:        "This event will start in 1 hour and last 30 minutes",
			ScheduledStartTime: &startingTime,
			ScheduledEndTime:   &endingTime,
			EntityType:         discordgo.GuildScheduledEventEntityTypeStageInstance,
			ChannelID:          "1061322715086205019",
			PrivacyLevel:       discordgo.GuildScheduledEventPrivacyLevelGuildOnly,
		})

		if err != nil {
			log.Printf("Error creating scheduled event: %v", err)
			return
		}

		fmt.Println("Created scheduled event:", scheduledEvent.Name)

		s.StageInstanceEdit("Test-Stage", &discordgo.StageInstanceParams{
			Topic: "New amazing topic set on time " + time.Now().String(),
		})
	}
}

func RunIt() {
	err := ReadConfig()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	Start()

	return
}
