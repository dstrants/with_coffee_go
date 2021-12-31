package covid

type CovidApiResponse struct {
	DataProvider []CovidCases `json:"dataProvider"`
}

// Model for covid cases per country
type CovidCases struct {
	DateStamp    string `json:"date_stamp" bson:"date_stamp"`
	CntConfirmed int    `json:"cnt_confirmed" bson:"cnt_confirmed"`
	CntDeath     int    `json:"cnt_death" bson:"cnt_death"`
	CntRecovered int    `json:"cnt_recovered" bson:"cnt_recovered"`
	CntActive    int    `json:"cnt_active" bson:"cnt_active"`
}

type CovidCasesModel struct {
	Date           string `bson:"date"`
	Confirmed      int    `bson:"confirmed"`
	Deaths         int    `bson:"deaths"`
	Recovered      int    `bson:"recovered"`
	Active         int    `bson:"active"`
	Country        string `bson:"country"`
	TodayConfirmed int    `bson:"today_confirmed"`
	TodayRecovered int    `bson:"today_recovered"`
	TodayActive    int    `bson:"today_active"`
	TodayDeaths    int    `bson:"today_deaths"`
}
