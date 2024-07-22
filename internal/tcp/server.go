package tcp

import (
	"fmt"
	"net"

	"github.com/gijsdb/go-mem-db/internal/memdb"
	"github.com/rs/zerolog"
)

type Server struct {
	logger   zerolog.Logger
	address  string
	listener net.Listener
	DB       memdb.DBI
	commands chan Command
}

func NewServer(address string, db memdb.DBI, logger zerolog.Logger) Server {
	return Server{
		logger:   logger,
		address:  address,
		DB:       db,
		commands: make(chan Command),
	}
}

func (s *Server) Start() {
	listener, err := net.Listen("tcp", s.address)
	if err != nil {
		panic(err)
	}
	defer listener.Close()
	s.listener = listener
	s.logger.Info().Msg(fmt.Sprintf("TCP server started on %s", s.address))
	s.HandleConnections()
}

func (s *Server) HandleConnections() {
	for {
		conn, err := s.listener.Accept()
		if err != nil {
			s.logger.Fatal().Msg(fmt.Sprintf("tcp server unable to accept connection: %v", err))
			continue
		}
		go s.ReadCommand(conn)
		go s.HandleCommand()
	}
}
