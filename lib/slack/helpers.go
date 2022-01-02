package slack

import (
	"fmt"
	"with_coffee/lib/hackernews"

	"github.com/slack-go/slack"
)

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

// Helper that appends all hackenews stories to the parent message as a separate block.
func HackeNewsMessageBlock(stories []hackernews.Story, message slack.Message) slack.Message {
	headerText := slack.NewTextBlockObject("mrkdwn", ":newspaper: *HACKER NEWS STORIES* :newspaper:", false, false)
	headerSection := slack.NewSectionBlock(headerText, nil, nil)

	message.Blocks.BlockSet = append(message.Blocks.BlockSet, headerSection)

	for i, story := range stories {
		msg := fmt.Sprintf("%d. <%s|%s> \n\t :man-pouting: _%s_ - :star-struck: %d", i+1, story.URL, story.Title, story.By, story.Score)
		message.Blocks.BlockSet = append(message.Blocks.BlockSet, slack.NewSectionBlock(slack.NewTextBlockObject("mrkdwn", msg, false, false), nil, nil))
	}
	message.Blocks.BlockSet = append(message.Blocks.BlockSet, slack.NewDividerBlock())

	return message
}
