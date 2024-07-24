package main

import (
	"fmt"

	"github.com/gijsdb/go-mem-db/internal/adapter/controller"
	"github.com/gijsdb/go-mem-db/pkg/config"
	"github.com/gijsdb/go-mem-db/pkg/logging"
	"github.com/spf13/viper"
)

func main() {
	logger := logging.CreateOrGetMultiOutputLogger()

	err := config.InitConfig("./")
	if err != nil {
		logger.Fatal().Msg("Failed to read config file")
	}

	TCP_ADDRESS := viper.GetString("tcp.address")
	TCP_PORT := viper.GetString("tcp.port")

	controller := controller.NewCLIController(logger)
	controller.HandleStartCLI(fmt.Sprintf("%s:%s", TCP_ADDRESS, TCP_PORT))
}
