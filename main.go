package main

import (
	"with_coffee/lib/covid"
	"with_coffee/lib/hackernews"
	"with_coffee/lib/slack"
	"with_coffee/lib/weather"
)

const Version = "0.3.2"

func main() {
	// Initialized slack message
	message := slack.InitWithCoffeeMessage()

	// Weather Forecast
	weather.StoreForecastToMongo()
	forecasts := weather.GetAllCitiesLocations()
	message = slack.WeatherForecastMessageBlock(forecasts, message)

	// Covid Stats
	covid.ImportAllCountriesCases()
	covidCases := covid.LoadCovidCases()
	message = slack.CovidMessageBlock(covidCases, message)

	// HackerNews Stories
	stories := hackernews.ImportStories()
	message = slack.HackerNewsMessageBlock(stories, message)

	// Post full message to slack
	slack.SendMultiBlockMessage(message)

}
