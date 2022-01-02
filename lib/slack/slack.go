package slack

import (
	"fmt"
	"log"

	"github.com/slack-go/slack"

	"with_coffee/lib/config"
)

// Sends a slack plain text message
func SendSimpleMessage(msg string) {
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

// Sends a message from markdown source
func SendMarkdownMessage(msg string) {
	cnf, _ := config.LoadConfig()
	api := slack.New(cnf.Slack.Token)

	_, _, err := api.PostMessage(
		cnf.Slack.Channel,
		slack.MsgOptionBlocks(slack.NewSectionBlock(slack.NewTextBlockObject("mrkdwn", msg, false, false), nil, nil)),
		slack.MsgOptionAsUser(true),
	)

	if err != nil {
		log.Fatal(err)
	}
}

// Send a a message from a slack.Message instance
func SendMultiBlockMessage(message slack.Message) {
	cnf, _ := config.LoadConfig()
	api := slack.New(cnf.Slack.Token)

	_, _, err := api.PostMessage(
		cnf.Slack.Channel,
		slack.MsgOptionBlocks(message.Blocks.BlockSet...),
		slack.MsgOptionAsUser(true),
	)

	if err != nil {
		log.Fatal(err)
	}

}
