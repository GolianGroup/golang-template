package configfx

import (
	"github.com/spf13/viper"
	"go.uber.org/fx"
	"log"
	"master/internal/pkg/config"
)

var Module = fx.Provide(initConfig)

func initConfig() *config.Application {
	viper.AutomaticEnv()
	viper.SetConfigFile("config/config.yml")
	viper.SetConfigType("yml")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("read config file failed, err:%v\n", err)
	}

	var conf config.Application

	if err := viper.Unmarshal(&conf); err != nil {
		log.Fatalf("unmarshal config file failed, err:%v\n", err)
	}

	return &conf
}
