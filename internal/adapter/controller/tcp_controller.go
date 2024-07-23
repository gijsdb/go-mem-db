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
	db      *memdb.DB
}

func NewTCPController(logger zerolog.Logger, address string, port string, db *memdb.DB) TCPController {
	return TCPController{
		logger:  logger,
		address: address,
		port:    port,
		db:      db,
	}
}

func (c *TCPController) HandleStartTCPServer() {
	s := tcp.NewServer(fmt.Sprintf("%s:%s", c.address, c.port), c.db, c.logger)
	s.Start()
}
