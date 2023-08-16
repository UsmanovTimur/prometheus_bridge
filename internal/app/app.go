package app

import (
	"flag"
	"fmt"
	"prometheus_bridge/internal/httpserver"
	"prometheus_bridge/internal/logging"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// Старт приложения
func Init() {
	logging.Init()
	//Чтение флагов
	host := flag.String("host", "127.0.0.1", "Server host. Default localhost")
	port := flag.String("port", "9091", "Server port. Default 9091")
	log_level := flag.Int("log-level", 4, "Log level 4- Info 5- Debug")
	flag.Parse()
	//чтение логов
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		logging.Logger.Fatal(fmt.Errorf("fatal error config file: %w", err))
	}
	logging.Logger.SetLevel(logrus.Level(*log_level))
	httpserver.Start(*host, *port)
}
