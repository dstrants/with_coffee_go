package covid

// Model for covid cases per country
type CovidCases struct {
	Country             string `json:"country" bson:"country"`
	Cases               int    `json:"cases" bson:"cases"`
	TodayCases          int    `json:"todayCases" bson:"todayCases"`
	Deaths              int    `json:"deaths" bson:"deaths"`
	TodayDeaths         int    `json:"todayDeaths" bson:"todayDeaths"`
	Recovered           int    `json:"recovered" bson:"recovered"`
	Active              int    `json:"active" bson:"active"`
	Critical            int    `json:"critical" bson:"critical"`
	CasesPerOneMillion  int    `json:"casesPerOneMillion" bson:"casesPerOneMillion"`
	DeathsPerOneMillion int    `json:"deathsPerOneMillion" bson:"deathsPerOneMillion"`
	TotalTests          int    `json:"totalTests" bson:"totalTests"`
	TestsPerOneMillion  int    `json:"testsPerOneMillion" bson:"testsPerOneMillion"`
	Date                string `bson:"date"`
}
