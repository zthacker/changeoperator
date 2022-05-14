package coSlackBot

import (
	log "github.com/sirupsen/logrus"
	"github.com/slack-go/slack"
)

func GenerateChangeModal(s slack.SlashCommand) {
	modalRequest := generateChangeModal(s)
	_, err := slackClient.OpenView(s.TriggerID, modalRequest)
	if err != nil {
		//this will catch the Slack-specific error on OpenView
		log.Errorf("Error on OpenView: %s", err)
	}
}
