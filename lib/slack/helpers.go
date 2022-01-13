package slack

import (
	"fmt"
	"with_coffee/lib/hackernews"

	"github.com/slack-go/slack"
)

// Instantiates a new Message that will send all info
func InitWithCoffeeMessage() slack.Message {
	header := slack.NewHeaderBlock(slack.NewTextBlockObject("plain_text", ":wave: :coffee: :sunny: Good Morning! Find your daily news below", false, false))
	message := slack.NewBlockMessage(header)
	message.Blocks.BlockSet = append(message.Blocks.BlockSet, slack.NewDividerBlock())
	return message
}

// Helper that appends all covid data to the parent message as a separate block.
func CovidMessageBlock(blocks []string, message slack.Message) slack.Message {
	headerText := slack.NewTextBlockObject("mrkdwn", ":microbe: *COVID CASES STATS* :microbe:", false, false)
	headerSection := slack.NewSectionBlock(headerText, nil, nil)

	message.Blocks.BlockSet = append(message.Blocks.BlockSet, headerSection)

	fields := make([]*slack.TextBlockObject, 0)
	for _, msg := range blocks {
		fields = append(fields, slack.NewTextBlockObject("mrkdwn", msg, false, false))
	}

	message.Blocks.BlockSet = append(message.Blocks.BlockSet, slack.NewSectionBlock(nil, fields, nil))
	message.Blocks.BlockSet = append(message.Blocks.BlockSet, slack.NewDividerBlock())

	return message
}

// Helper that appends all hackenews stories to the parent message as a separate block.
func HackerNewsMessageBlock(stories []hackernews.Story, message slack.Message) slack.Message {
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

func WeatherForecastMessageBlock(citiesForectasts []string, message slack.Message) slack.Message {
	headerText := slack.NewTextBlockObject("mrkdwn", ":sun_with_face: *WEATHER FORECAST* :sun_with_face:", false, false)
	headerSection := slack.NewSectionBlock(headerText, nil, nil)
	message.Blocks.BlockSet = append(message.Blocks.BlockSet, headerSection)

	fields := make([]*slack.TextBlockObject, 0)
	for _, forecast := range citiesForectasts {
		fields = append(fields, slack.NewTextBlockObject("mrkdwn", forecast, false, false))
	}

	message.Blocks.BlockSet = append(message.Blocks.BlockSet, slack.NewSectionBlock(nil, fields, nil))
	message.Blocks.BlockSet = append(message.Blocks.BlockSet, slack.NewDividerBlock())
	return message
}
