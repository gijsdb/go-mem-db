package main

import (
	"github.com/gijsdb/go-mem-db/internal/adapter/controller"
	"github.com/gijsdb/go-mem-db/pkg/logging"
)

func main() {
	logger := logging.CreateOrGetMultiOutputLogger()

	controller := controller.NewController(logger)
	controller.HandleStartTCPServer()
}
