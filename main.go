package main

import (
	"with_coffee/lib/covid"
	"with_coffee/lib/slack"
)

func main() {
	covid.ImportAllCountriesCases()

	slack.SendMarkdownMessage(covid.LoadCovidCases())
}
