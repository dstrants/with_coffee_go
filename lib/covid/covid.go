package covid

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/go-resty/resty/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"with_coffee/lib/config"
)

// Saves all covid data for today to the mongodb instance
func StoreCasesToMongo(cases []CovidCases) {
	cnf, _ := config.LoadConfig()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(cnf.Mongo.Uri))

	if err != nil {
		panic(err)
	}

	collection := client.Database(cnf.Mongo.Database).Collection("covid")

	documents := make([]interface{}, len(cases))

	for i, v := range cases {
		documents[i] = v
	}

	_, err = collection.InsertMany(ctx, documents)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Entries saved to mongo")
}

// Load cases for all countries for the third party covid api
func fetchCases() []CovidCases {
	var Results []CovidCases

	cnf, _ := config.LoadConfig()

	client := resty.New()
	_, err := client.R().SetResult(&Results).Get(cnf.Covid.Uri)

	if err != nil {
		log.Fatal(err)
	}
	return Results
}

// Load covid data for a specific country from the mongodb
func fetchCountryCases(country string) CovidCases {
	var Results CovidCases
	cnf, _ := config.LoadConfig()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(cnf.Mongo.Uri))
	if err != nil {
		log.Fatal(err)
	}

	collection := client.Database(cnf.Mongo.Database).Collection("covid")

	err = collection.FindOne(ctx, bson.D{{"country", country}}).Decode(&Results)

	if err != nil {
		log.Fatal(err)
	}

	return Results
}

// Adds date field to all cases structs.
func addTimestampToCases(cases []CovidCases) []CovidCases {
	var timestampedResults []CovidCases
	now := time.Now().Format("2006-01-02")
	for _, c := range cases {
		c.Date = now
		timestampedResults = append(timestampedResults, c)
	}
	log.Printf("Added current timestamp to %d entries", len(timestampedResults))
	return timestampedResults
}

// Wrapper function to consume the covid api, add timestamps and save the covid data
// to mongodb
func ImportCovidCases() []CovidCases {
	log.Println("Starting importing of cases from the API...")
	cases := addTimestampToCases(fetchCases())

	StoreCasesToMongo(cases)

	return cases
}

// Prepare a message for covid status
func LoadCovidCases() string {
	results := fetchCountryCases("Greece")

	msg := fmt.Sprintf("%s\n * Cases | New: %v Total: %v \n * Deaths | New: %v Total: %v\n", results.Country, results.TodayCases, results.Cases, results.TodayDeaths, results.TodayDeaths)
	return msg
}
