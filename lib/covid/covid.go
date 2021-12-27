package covid

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/go-resty/resty/v2"
	"go.mongodb.org/mongo-driver/bson"

	"with_coffee/lib/config"
	db "with_coffee/lib/mongo"
)

// Checks if the sync has already been done for today
func NeedsToImport() bool {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := db.MongoCollection(ctx, "syncs")

	var check bson.D

	collection.FindOne(ctx, bson.D{{"type", "covid"}, {"date", time.Now().Format("2006-01-02")}}).Decode(&check)

	return check == nil
}

// Saves all covid data for today to the mongodb instance
// The function also checks if the sync has already been done for today
// And cancel the sync if it has.
func StoreCasesToMongo(cases []CovidCases) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := db.MongoCollection(ctx, "covid")

	documents := make([]interface{}, len(cases))

	for i, v := range cases {
		documents[i] = v
	}

	_, err := collection.InsertMany(ctx, documents)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Entries saved to mongo")

	collection = db.MongoCollection(ctx, "syncs")

	_, err = collection.InsertOne(ctx, bson.D{{"type", "covid"}, {"date", time.Now().Format("2006-01-02")}})

	log.Println("Marked sync as successful")
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

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := db.MongoCollection(ctx, "covid")

	err := collection.FindOne(ctx, bson.D{{"country", country}}).Decode(&Results)

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
func ImportCovidCases() {
	if !NeedsToImport() {
		log.Printf("Imported has been performed already for %s, skipping...\n", time.Now().Format("2006-01-02"))
		return
	}

	log.Println("Starting importing of cases from the API...")
	cases := addTimestampToCases(fetchCases())

	StoreCasesToMongo(cases)
}

// Prepare a message for covid status
func LoadCovidCases() string {
	cnf, _ := config.LoadConfig()
	countries := strings.Split(cnf.Covid.Countries, ",")

	var msg string

	for _, country := range countries {
		results := fetchCountryCases(country)

		msg = msg + fmt.Sprintf("%s\n * Cases | New: %v Total: %v \n * Deaths | New: %v Total: %v\n", results.Country, results.TodayCases, results.Cases, results.TodayDeaths, results.TodayDeaths)
	}

	return msg
}
