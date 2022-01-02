package main

import (
	"with_coffee/lib/hackernews"
	"with_coffee/lib/slack"
)

const Version = "0.2.4"

func main() {
	// covid.ImportAllCountriesCases()
	message := slack.InitWithCoffeeMessage()

	// // Covid Stats
	// covidHeader, covidStats := slack.CovidMessageBlock(covid.LoadCovidCases())
	// message.Blocks.BlockSet = append(message.Blocks.BlockSet, covidHeader)
	// message.Blocks.BlockSet = append(message.Blocks.BlockSet, covidStats)
	// message.Blocks.BlockSet = append(message.Blocks.BlockSet, slackApi.NewDividerBlock())

	stories := hackernews.ImportStories()

	message = slack.HackeNewsMessageBlock(stories, message)

	slack.SendMultiBlockMessage(message)
}
