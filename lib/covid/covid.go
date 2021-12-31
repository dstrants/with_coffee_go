package covid

import (
	"context"
	"fmt"
	"log"
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

// Adds extra fields to cases structs
func addCountryToCases(country string, cases []CovidCases) []CovidCasesModel {
	documents := make([]CovidCasesModel, 0)

	for i, v := range cases {
		var model CovidCasesModel

		model.Country = country
		model.Active = v.CntActive
		model.Confirmed = v.CntConfirmed
		model.Deaths = v.CntDeath
		model.Active = v.CntActive
		model.Date = v.DateStamp

		if i > 0 {
			model.TodayActive = v.CntActive - cases[i-1].CntActive
			model.TodayConfirmed = v.CntConfirmed - cases[i-1].CntConfirmed
			model.TodayRecovered = v.CntRecovered - cases[i-1].CntRecovered
			model.TodayDeaths = v.CntDeath - cases[i-1].CntDeath
		} else {
			model.TodayActive = 0
			model.TodayConfirmed = 0
			model.TodayRecovered = 0
			model.TodayDeaths = 0
		}
		documents = append(documents, model)

	}
	return documents
}

// Saves all covid data for today to the mongodb instance
// The function also checks if the sync has already been done for today
// And cancel the sync if it has.
func StoreCasesToMongo(cases []CovidCasesModel) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	latestDate := cases[len(cases)-1].Date

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

	collection.FindOneAndUpdate(ctx, bson.D{{"type", "covid"}}, bson.D{{"date", latestDate}})

	log.Println("Marked sync as successful")
}

// Load cases for all countries for the third party covid api
func fetchCases(country string) CovidApiResponse {
	var Results CovidApiResponse

	cnf, _ := config.LoadConfig()

	client := resty.New()
	_, err := client.R().
		SetQueryParams(map[string]string{
			"cols":   "date_stamp,cnt_confirmed,cnt_death,cnt_recovered,cnt_active",
			"where":  fmt.Sprintf("(iso3166_1=%s)", country),
			"format": "amcharts",
			"limit":  "5000",
		}).
		SetResult(&Results).Get(cnf.Covid.Uri)

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

	fmt.Println(Results)
	return Results
}

// Wrapper function to consume the covid api, add timestamps and save the covid data
// to mongodb
func ImportCovidCases(country string) {
	if !NeedsToImport() {
		log.Printf("Imported has been performed already for %s, skipping...\n", time.Now().Format("2006-01-02"))
		return
	}

	log.Println("Starting importing of cases from the API...")
	cases := addCountryToCases(country, fetchCases(country).DataProvider)

	StoreCasesToMongo(cases)
}

// Prepare a message for covid status
// func LoadCovidCases() string {
// 	cnf, _ := config.LoadConfig()
// 	countries := strings.Split(cnf.Covid.Countries, ",")

// 	var msg string

// 	// for _, country := range countries {
// 	// 	results := fetchCountryCases(country)

// 	// 	//msg = msg + fmt.Sprintf("%s\n * Cases: %v Total: %v \n * Deaths | New: %v Total: %v\n", results.Country, results.CN, results.Cases, results.TodayDeaths, results.TodayDeaths)
// 	// }

// 	return msg
// }
