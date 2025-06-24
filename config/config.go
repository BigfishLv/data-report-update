package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	Env    string
	MySQL  mysql
	Logger logger
	Csv    csv
}

type logger struct {
	Path    string
	Level   string
	Console bool
}

type mysql struct {
	Host     string
	Port     string
	Username string
	Password string
	Database string
	Charset  string
	TimeZone string
}

type csv struct {
	Path                                  string
	AdPvClickCostDailyCsvFileName         string
	CampaignDataDailySummariesCsvFileName string
	CreativeDataDailySummariesCsvFileName string
	UserDataDailySummariesCsvFileName     string
	AllUsersDataDailySummariesCsvFileName string
}

func NewConfig(params *Params) *Config {
	conf := Config{}
	viper.SetConfigName("config" + "_" + params.Env)
	viper.SetConfigType("yaml")
	viper.AddConfigPath(params.ConfigPath)

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	if err := viper.Unmarshal(&conf); err != nil {
		panic(fmt.Errorf("unable to decode config: %s", err))
	}

	return &conf
}
