package main

import (
	"flag"

	"github.com/gijsdb/go-mem-db/internal/adapter/controller"
	"github.com/gijsdb/go-mem-db/pkg/logging"
)

func main() {
	logger := logging.CreateOrGetMultiOutputLogger()

	TCP_ADDRESS := flag.String("tcp-address", "localhost", "TCP server address, defaults to localhost")
	TCP_PORT := flag.String("tcp-port", "4242", "TCP server port, defaults to 4242")

	flag.Parse()

	controller := controller.NewTCPController(logger, *TCP_ADDRESS, *TCP_PORT)
	controller.HandleStartTCPServer()
}
