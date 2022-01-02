package main

import (
	"with_coffee/lib/covid"
	"with_coffee/lib/hackernews"
	"with_coffee/lib/slack"
)

const Version = "0.3.0"

func main() {
	// Initialized slack message
	message := slack.InitWithCoffeeMessage()

	// Covid Stats
	covid.ImportAllCountriesCases()
	covidCases := covid.LoadCovidCases()
	message = slack.CovidMessageBlock(covidCases, message)

	// HackerNews Stories
	stories := hackernews.ImportStories()
	message = slack.HackerNewsMessageBlock(stories, message)

	slack.SendMultiBlockMessage(message)
}
