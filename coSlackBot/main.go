package main

import (
	coSlackBot "coSlackBot/internal"
	log "github.com/sirupsen/logrus"
	"github.com/slack-go/slack"
	"github.com/slack-go/slack/slackevents"
	"net/http"
)

func main() {

	log.Info("Starting up ChangeBot...")
	eventsChannel := make(chan slackevents.EventsAPIInnerEvent)
	commandsChannel := make(chan slack.SlashCommand)
	interactiveChannel := make(chan slack.InteractionCallback)

	log.Info("Setting channels")
	coSlackBot.EventChannel = eventsChannel
	coSlackBot.CommandChannel = commandsChannel
	coSlackBot.InteractiveChannel = interactiveChannel

	log.Info("Running Inits")
	go coSlackBot.RunSlackInteractive()
	go coSlackBot.RunSlackCommands()

	log.Info("Handle-funcs-a-go-go")
	http.HandleFunc("/events", coSlackBot.EventHandler)
	http.HandleFunc("/command", coSlackBot.CommandHandler)
	http.HandleFunc("/interactive-endpoint", coSlackBot.InteractiveHandler)

	//TODO env var the port
	log.Info("Listen and Serve")
	err := http.ListenAndServe(":9500", nil)
	if err != nil {
		log.Fatalf(err.Error())
	}
}
