package main

import (
	"github.com/gijsdb/go-mem-db/internal/adapter/controller"
	"github.com/gijsdb/go-mem-db/pkg/logging"
)

func main() {
	logger := logging.CreateOrGetMultiOutputLogger()

	// TCP_ADDRESS := flag.String("tcp-address", "localhost:4242", "TCP server address, defaults to localhost:4242")

	controller := controller.NewCLIController(logger)
	controller.HandleStartCLI()
}
