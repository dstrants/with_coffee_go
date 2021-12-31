package covid

type CovidApiResponse struct {
	DataProvider []CovidCases `json:"dataProvider"`
}

// Model for covid cases per country
type CovidCases struct {
	DateStamp    string `json:"date_stamp" bson:"date_stamp"`
	CntConfirmed int64  `json:"cnt_confirmed" bson:"cnt_confirmed"`
	CntDeath     int64  `json:"cnt_death" bson:"cnt_death"`
	CntRecovered int64  `json:"cnt_recovered" bson:"cnt_recovered"`
	CntActive    int64  `json:"cnt_active" bson:"cnt_active"`
}

type CovidSyncModel struct {
	Date     string `bson:"date"`
	Metadata struct {
		Country string `bson:"country"`
	} `bson:"metadata"`
	Type string `bson:"type"`
}

type CovidCasesModel struct {
	Date           string `bson:"date"`
	Confirmed      int64  `bson:"confirmed"`
	Deaths         int64  `bson:"deaths"`
	Recovered      int64  `bson:"recovered"`
	Active         int64  `bson:"active"`
	Country        string `bson:"country"`
	TodayConfirmed int64  `bson:"today_confirmed"`
	TodayRecovered int64  `bson:"today_recovered"`
	TodayActive    int64  `bson:"today_active"`
	TodayDeaths    int64  `bson:"today_deaths"`
}
