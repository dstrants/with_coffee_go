package weather

import (
	"bytes"
	"fmt"
	"log"
	"strings"
	"text/template"
	"with_coffee/lib/config"

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
	cities := strings.Split(cnf.Weather.Locations, ",")

	paths := []string{"./lib/weather/message.tpl"}

	message, err := template.New("message.tpl").ParseFiles(paths...)
	if err != nil {
		log.Println("Could not generate template")
		panic(err)
	}

	var msg []string
	for _, city := range cities {
		forecast := GetLocationForecast(city)

		var tpl bytes.Buffer
		err = message.Execute(&tpl, forecast)

		if err != nil {
			panic(err)
			log.Println("Could not render template")
		}

		msg = append(msg, tpl.String())

	}
	return msg
}
