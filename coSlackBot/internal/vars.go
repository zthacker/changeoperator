package coSlackBot

import (
	"github.com/slack-go/slack"
	"github.com/slack-go/slack/slackevents"
	"os"
)

var EventChannel chan slackevents.EventsAPIInnerEvent
var CommandChannel chan slack.SlashCommand
var InteractiveChannel chan slack.InteractionCallback
var slackClient = slack.New(os.Getenv("SLACK_TOKEN"))

type ChangeRequest struct {
	Requester      string `json:"requester"`
	Env            string `json:"env"`
	Type           string `json:"type"`
	CustomerImpact string `json:"customer_impact"`
	Description    string `json:"description"`
	Date           string `json:"date"`
	Link           string `json:"link"`
	LinkBack       string `json:"link_back"`
}
