package slack

import (
	"fmt"

	"github.com/slack-go/slack"

	"with_coffee/lib/config"
)

// Sends a slack message
func Send(msg string) {
	cnf, _ := config.LoadConfig()

	api := slack.New(cnf.Slack.Token)

	channelID, timestamp, err := api.PostMessage(
		cnf.Slack.Channel,
		slack.MsgOptionText(msg, false),
		slack.MsgOptionAsUser(true),
	)
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
	fmt.Printf("Message successfully sent to channel %s at %s", channelID, timestamp)
}
