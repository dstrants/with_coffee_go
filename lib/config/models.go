package config

// The model that the environment configuration is parsed to
type Config struct {
	Covid struct {
		Uri string `env:"COVID_API_BASE_URL,default=https://coronavirus-19-api.herokuapp.com/countries"`
	}

	Mongo struct {
		Uri      string `env:"MONGO_CONNECTION_STRING"`
		Database string `env:"MONGO_DATABASE,default=news"`
	}

	Slack struct {
		Token   string `env:"SLACK_TOKEN"`
		Channel string `env:"SLACK_CHANNEL,default=news"`
	}
}
