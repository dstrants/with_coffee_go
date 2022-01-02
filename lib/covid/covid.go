package covid

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/dustin/go-humanize"
	"github.com/go-resty/resty/v2"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"

	"with_coffee/lib/config"
	db "with_coffee/lib/mongo"
)

// Checks if the sync has already been done for today
func NeedsToImport(country string) CovidSyncModel {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var filter CovidSyncModel

	filter.Metadata.Country = country
	filter.Type = "covid"

	collection := db.MongoCollection(ctx, "syncs")

	var check CovidSyncModel

	collection.FindOne(ctx, bson.D{{"type", "covid"}, {"metadata.country", country}}).Decode(&check)

	if check.Date == "" {
		log.Println("This is the first covid sync. Creating a placeholder entry on sync table.")
		filter.Date = "1970-01-01"
		collection.InsertOne(ctx, filter)
	}
	filter.Date = ""
	collection.FindOne(ctx, filter).Decode(&check)

	return check
}

// Adds extra fields to cases structs
func addCountryToCases(country string, cases []CovidCasesResponse) []CovidCasesModel {
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

	if len(cases) < 1 {
		log.Println("No data found. Skipping sync")
		return
	}

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

	log.Printf("%d entries saved to mongo", len(documents))

	collection = db.MongoCollection(ctx, "syncs")

	collection.FindOneAndUpdate(ctx, bson.D{{"type", "covid"}, {"metadata.country", cases[0].Country}}, bson.D{{"$set", bson.D{{"date", latestDate}}}})

	log.Printf("Marked sync as successful until %s\n", latestDate)
}

// Load cases for all countries for the third party covid api
func fetchCases(country string, date string) CovidApiResponse {
	var Results CovidApiResponse

	filter := fmt.Sprintf("(iso3166_1=%s)", country)

	if date != "" {
		filter = filter + " AND " + fmt.Sprintf("(date_stamp>%s)", date)
	}

	cnf, _ := config.LoadConfig()

	client := resty.New()
	_, err := client.R().
		SetQueryParams(map[string]string{
			"cols":   "date_stamp,cnt_confirmed,cnt_death,cnt_recovered,cnt_active",
			"where":  filter,
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
func fetchCountryCases(country string) CovidCasesModel {
	var Results CovidCasesModel

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := db.MongoCollection(ctx, "covid")

	var filters options.FindOneOptions

	filters.Sort = bson.D{{"date", -1}}

	err := collection.FindOne(ctx, bson.D{{"country", country}}, &filters).Decode(&Results)

	if err != nil {
		log.Fatal(err)
	}

	return Results
}

// Wrapper function to consume the covid api, add timestamps and save the covid data
// to mongodb
func ImportCountryCases(country string) {
	syncProps := NeedsToImport(country)
	log.Printf("Starting importing of cases for %s after %s from the API...\n", country, syncProps.Date)
	cases := addCountryToCases(country, fetchCases(country, syncProps.Date).DataProvider)

	StoreCasesToMongo(cases)
}

// Import cases for all countries configured
func ImportAllCountriesCases() {
	cnf, _ := config.LoadConfig()
	countries := strings.Split(cnf.Covid.Countries, ",")

	for _, country := range countries {
		ImportCountryCases(country)
	}
}

// Prepare a message for covid status
func LoadCovidCases() string {
	cnf, _ := config.LoadConfig()
	countries := strings.Split(cnf.Covid.Countries, ",")

	var msg string

	for _, country := range countries {
		results := fetchCountryCases(country)

		msg = msg + fmt.Sprintf(":flag-%s:  *%s*\n • Cases\n\t○ New → `%s`\n\t○ Total → `%s` \n • Deaths\n\t○ New → `%s`\n\t○ Total → `%s`\n\n",
			strings.ToLower(results.Country),
			results.Country,
			humanize.Comma(results.TodayConfirmed),
			humanize.Comma(results.Confirmed),
			humanize.Comma(results.TodayDeaths),
			humanize.Comma(results.Deaths),
		)
	}
	return msg
}
