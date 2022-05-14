package coSlackBot

import (
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	log "github.com/sirupsen/logrus"
	"github.com/slack-go/slack"
	"os"
	"strconv"
	"strings"
	"time"
)

func HandleCreateChange(callback slack.InteractionCallback) {
	r := resty.New()

	//setup the change request
	cr := makeChangeRequest(callback)

	//post to channel to get the Slack-specific message link
	channelID, respTimeStamp, toChannelErr := postToChangeChannel(*cr)
	if toChannelErr != nil {
		log.Errorf("Error posting to Change Channel: %s", toChannelErr)
		slackClient.PostMessage(callback.User.ID, slack.MsgOptionText("Sorry, there was an error in making your change request; logs will provide more details", false))
		return
	}

	//curate and set the linkback
	linkBack := curateLinkBack(channelID, respTimeStamp)
	cr.LinkBack = linkBack

	//marshal the json
	body, err := json.Marshal(cr)
	if err != nil {
		log.Error(err)
		slackClient.PostMessage(callback.User.ID, slack.MsgOptionText("Sorry, there was an error in making your change request; logs will provide more details", false))
		return
	}

	//POST to changeoperator
	postChangeErr := postToChangeOperator(r, body)
	if postChangeErr != nil {
		log.Errorf("Erroring POSTing to ChangeOperator: %s", postChangeErr)
		slackClient.PostMessage(callback.User.ID, slack.MsgOptionText("Sorry, there was an error in making your change request; logs will provide more details", false))
		return
	}

	//let the requesting Slack user know what happened
	msgToUser := fmt.Sprintf("Thanks for creating a change! You can view it here: %s", cr.LinkBack)
	_, _, postMsgErr := slackClient.PostMessage(callback.User.ID, slack.MsgOptionText(msgToUser, false))
	if postMsgErr != nil {
		log.Errorf("Couldn't post message to User: %s, because of: %s", callback.User.Name, postMsgErr)
		return
	}

}

func makeChangeRequest(i slack.InteractionCallback) *ChangeRequest {
	cr := ChangeRequest{
		Requester:      i.User.Name,
		Env:            i.View.State.Values["Environment"]["environment"].SelectedOption.Value,
		Type:           i.View.State.Values["Changetype"]["changetype"].SelectedOption.Value,
		CustomerImpact: i.View.State.Values["CustomerImpact"]["customerimpact"].SelectedOption.Value,
		Description:    i.View.State.Values["Description"]["description"].Value,
		Date:           time.Now().Format("01-02-2006"),
		Link:           i.View.State.Values["Link"]["link"].Value,
	}

	return &cr
}

func postToChangeChannel(cr ChangeRequest) (string, string, error) {

	msg := fmt.Sprintf("*Change Created by:* %s\n*Env*: %s\n*Type:* %s\n*Customer Impact?:* %s\n*Description:* %s\n*Change Date:* %s\nLink: %s", cr.Requester, cr.Env, cr.Type, cr.CustomerImpact, cr.Description, cr.Date, cr.Link)
	cID, rTS, err := slackClient.PostMessage(os.Getenv("CHANGE_CHANNEL_ID"), slack.MsgOptionText(msg, false))
	if err != nil {
		log.Error(err)
		return "", "", err
	}
	return cID, rTS, nil
}

func curateLinkBack(channelID, responseTimeStamp string) string {
	joinMsg := strings.Join(strings.Split(responseTimeStamp, "."), "")
	dMsg, err := strconv.Atoi(joinMsg)
	if err != nil {
		log.Errorf("Error converting timestamp: %s", err)
		return ""
	}
	msgLink := fmt.Sprintf("%s/archives/%s/p%d", os.Getenv("SLACK_WORKSPACE_URL"), channelID, dMsg)

	return msgLink
}

func postToChangeOperator(r *resty.Client, body []byte) error {
	token := os.Getenv("CHANGE_OPERATOR_TOKEN")
	_, err := r.R().
		SetHeader("Accept", "application/json; indent=4").
		SetHeader("api-key", token).
		SetBody(body).Post(os.Getenv("CHANAGE_OPERATOR_URL"))
	if err != nil {
		return err
	}
	return nil
}
