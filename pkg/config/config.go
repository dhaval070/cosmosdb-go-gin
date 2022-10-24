package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	DbKey              string `mapstructure:"DB_KEY"`
	DbEndpoint         string `mapstructure:"DB_ENDPOINT"`
	Address            string `mapstructure:"ADDRESS"`
	InstrumentationKey string `mapstructure:"INSTRUMENTATION_KEY"`
}

var vars = []string{"DB_KEY", "DB_ENDPOINT", "INSTRUMENTATION_KEY"}

func Load() *Config {
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.SetEnvPrefix("COSMOSAPP")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetDefault("address", ":8080")

	for _, v := range vars {
		if err := viper.BindEnv(v); err != nil {
			log.Fatal(err)
		}
	}

	if err := viper.ReadInConfig(); err != nil {
		log.Println(err)
	}

	cfg := Config{}
	if err := viper.Unmarshal(&cfg); err != nil {
		log.Fatal(err)
	}

	return &cfg
}
