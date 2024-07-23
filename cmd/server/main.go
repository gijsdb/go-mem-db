package main

import (
	"flag"

	"github.com/gijsdb/go-mem-db/internal/adapter/controller"
	"github.com/gijsdb/go-mem-db/internal/memdb"
	"github.com/gijsdb/go-mem-db/pkg/logging"
)

func main() {
	logger := logging.CreateOrGetMultiOutputLogger()

	TCP_ADDRESS := flag.String("tcp-address", "localhost", "TCP server address, defaults to localhost")
	TCP_PORT := flag.String("tcp-port", "4242", "TCP server port, defaults to 4242")
	WEB_UI_PORT := flag.String("web-port", "4141", "Web UI port, defaults to 4141")
	flag.Parse()

	db := memdb.NewDB(logger)

	tcp_controller := controller.NewTCPController(logger, *TCP_ADDRESS, *TCP_PORT, db)
	go tcp_controller.HandleStartTCPServer()

	web_ui_controller := controller.NewWebUIController(logger, *WEB_UI_PORT, db)
	web_ui_controller.HandleStart()
}
