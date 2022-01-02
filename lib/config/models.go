package config

// The model that the environment configuration is parsed to
type Config struct {
	Covid struct {
		Uri       string `env:"COVID_API_BASE_URL,default=https://covid19.richdataservices.com/rds/api/query/int/jhu_country/select"`
		Countries string `env:"COVID_COUNTRIES,default=GR"`
	}

	Mongo struct {
		Uri      string `env:"MONGO_CONNECTION_STRING,required=true"`
		Database string `env:"MONGO_DATABASE,default=news"`
	}

	Slack struct {
		Token   string `env:"SLACK_TOKEN,required=true"`
		Channel string `env:"SLACK_CHANNEL,default=news"`
	}

	HackerNews struct {
		BaseUrl string `env:"HACKERNEWS_BASE_URL,default=https://hacker-news.firebaseio.com/v0"`
		Limit   int    `env:"HACKERNEWS_STORIES_LIMIT,default=5"`
	}
}
