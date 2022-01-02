package slack

import "github.com/slack-go/slack"

// Instantiates a new Message that will send all info
func InitWithCoffeeMessage() slack.Message {
	header := slack.NewHeaderBlock(slack.NewTextBlockObject("plain_text", ":wave: :coffee: :sunny: Good Morning! Find your daily news below", false, false))
	return slack.NewBlockMessage(header)
}

// Helper that appends all covid data to the parent message as a separate block.
func CovidMessageBlock(blocks []string) (*slack.SectionBlock, *slack.SectionBlock) {
	headerText := slack.NewTextBlockObject("mrkdwn", ":microbe: *Covid Cases Stats*", false, false)
	headerSection := slack.NewSectionBlock(headerText, nil, nil)

	fields := make([]*slack.TextBlockObject, 0)
	for _, msg := range blocks {
		fields = append(fields, slack.NewTextBlockObject("mrkdwn", msg, false, false))
	}

	return headerSection, slack.NewSectionBlock(nil, fields, nil)
}
