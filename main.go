package main

import (
	"with_coffee/lib/covid"
	"with_coffee/lib/slack"

	slackApi "github.com/slack-go/slack"
)

const Version = "0.2.3"

func main() {
	covid.ImportAllCountriesCases()
	message := slack.InitWithCoffeeMessage()

	// Covid Stats
	covidHeader, covidStats := slack.CovidMessageBlock(covid.LoadCovidCases())
	message.Blocks.BlockSet = append(message.Blocks.BlockSet, covidHeader)
	message.Blocks.BlockSet = append(message.Blocks.BlockSet, covidStats)
	message.Blocks.BlockSet = append(message.Blocks.BlockSet, slackApi.NewDividerBlock())

	slack.SendMultiBlockMessage(message)

}
