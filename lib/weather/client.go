package weather

import (
	"bytes"
	"fmt"
	"log"
	"strings"
	"with_coffee/lib/config"
	"with_coffee/lib/format"

	"github.com/go-resty/resty/v2"
)

var cnf, _ = config.LoadConfig()

// Loads weather forecast for a given location
func GetLocationForecast(location string) Weather {
	url := fmt.Sprintf("%s/forecast.json", cnf.Weather.BaseUrl)

	var weather Weather
	client := resty.New()

	client.R().
		SetQueryParams(map[string]string{
			"key": cnf.Weather.Token,
			"q":   location,
		}).
		SetResult(&weather).Get(url)

	return weather
}

// Loads weather for all locations
func GetAllCitiesLocations() []string {
	var msg []string

	cities := strings.Split(cnf.Weather.Locations, ",")
	message, err := format.LoadTemplate("message.tpl", []string{"./lib/format/templates/weather/message.tpl"})

	if err != nil {
		log.Printf("Could not generate template. Error: %v", err)
		return msg
	}

	for _, city := range cities {
		forecast := GetLocationForecast(city)

		var tpl bytes.Buffer
		err = message.Execute(&tpl, forecast)

		if err != nil {
			log.Printf("Could not render template. Error: %v", err)
		}

		msg = append(msg, tpl.String())

	}
	return msg
}
