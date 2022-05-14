package coSlackBot

import (
	"encoding/json"
	"fmt"
	"github.com/slack-go/slack"
	"github.com/slack-go/slack/slackevents"
	"io/ioutil"
	"net/http"
)

// Handler for Slack Events
func EventHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	eventsApiEvent, err := slackevents.ParseEvent(body, slackevents.OptionNoVerifyToken())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if eventsApiEvent.Type == slackevents.URLVerification {
		var r *slackevents.ChallengeResponse
		err := json.Unmarshal(body, &r)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "text")
		w.Write([]byte(r.Challenge))
	}

	if eventsApiEvent.Type == slackevents.CallbackEvent {
		EventChannel <- eventsApiEvent.InnerEvent
	}
}

// Handler for Slack Interactives such as modals, callbacks, button-clicks, etc
func InteractiveHandler(w http.ResponseWriter, r *http.Request) {
	var i = slack.InteractionCallback{}
	err := json.Unmarshal([]byte(r.FormValue("payload")), &i)
	if err != nil {
		fmt.Printf(err.Error())
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	InteractiveChannel <- i
}

func CommandHandler(w http.ResponseWriter, r *http.Request) {
	slashCommand, err := slack.SlashCommandParse(r)
	if err != nil {
		fmt.Printf(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	CommandChannel <- slashCommand
}
