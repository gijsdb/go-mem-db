package controller

import (
	"fmt"

	"github.com/gijsdb/go-mem-db/internal/memdb"
	"github.com/gijsdb/go-mem-db/internal/tcp"
	"github.com/rs/zerolog"
)

type TCPController struct {
	logger  zerolog.Logger
	address string
	port    string
}

func NewTCPController(logger zerolog.Logger, address string, port string) TCPController {
	return TCPController{
		logger:  logger,
		address: address,
		port:    port,
	}
}

func (c *TCPController) HandleStartTCPServer() {
	db := memdb.NewDB(c.logger)
	s := tcp.NewServer(fmt.Sprintf("%s:%s", c.address, c.port), db, c.logger)
	s.Start()
}
