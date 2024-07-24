package main

import (
	"github.com/gijsdb/go-mem-db/internal/adapter/controller"
	"github.com/gijsdb/go-mem-db/internal/memdb"
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

	WEB_UI_PORT := viper.GetString("web.port")
	RUN_WEB_UI := viper.GetBool("web.run")

	db := memdb.NewDB(logger)

	tcp_controller := controller.NewTCPController(logger, TCP_ADDRESS, TCP_PORT, db)

	if RUN_WEB_UI {
		go tcp_controller.HandleStartTCPServer()
		web_ui_controller := controller.NewWebUIController(logger, WEB_UI_PORT, db)
		web_ui_controller.HandleStart()
	} else {
		tcp_controller.HandleStartTCPServer()
	}
}
