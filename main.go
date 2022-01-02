package main

import (
	"with_coffee/lib/covid"
	"with_coffee/lib/slack"
)

const Version = "0.2.3"

func main() {
	covid.ImportAllCountriesCases()

	slack.SendMarkdownMessage(covid.LoadCovidCases())
}
