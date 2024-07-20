package controller

import (
	"github.com/gijsdb/go-mem-db/internal/memdb"
	"github.com/gijsdb/go-mem-db/internal/tcp"
	"github.com/rs/zerolog"
)

type Controller struct {
	logger zerolog.Logger
}

func NewController(logger zerolog.Logger) Controller {
	return Controller{
		logger: logger,
	}
}

func (c *Controller) HandleStartTCPServer() {
	db := memdb.NewDB(c.logger)
	s := tcp.NewServer("localhost:4242", db, c.logger)
	s.Start()
	c.logger.Info().Msg("TCP server started")
}
