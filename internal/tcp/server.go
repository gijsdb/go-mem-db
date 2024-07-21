package tcp

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"time"

	"github.com/gijsdb/go-mem-db/internal/memdb"
	"github.com/rs/zerolog"
)

type Command struct {
	Value string
	Args  []string
	Conn  net.Conn
}

type Server struct {
	logger   zerolog.Logger
	address  string
	listener net.Listener
	DB       memdb.DB
	commands chan Command
}

func NewServer(address string, db memdb.DB, logger zerolog.Logger) Server {
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

func (s *Server) ReadCommand(conn net.Conn) {
	defer conn.Close()
	for {
		input, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			s.logger.Info().Msg(fmt.Sprintf("no message received on connection... waiting: %v", err))
			time.Sleep(10 * time.Second)
		}
		cmd := strings.Trim(input, "\r\n")
		args := strings.Split(cmd, " ")

		s.commands <- Command{
			Value: args[0], Args: args[1:], Conn: conn,
		}
	}
}

func (s *Server) WriteCommand(conn net.Conn, data string) {
	if _, err := conn.Write([]byte(data + "\n")); err != nil {
		s.logger.Error().Msg(fmt.Sprintf("TCP server failed to write : %v", err))
	}
}

func (s *Server) HandleCommand() {
	for cmd := range s.commands {
		switch cmd.Value {
		case LIST:
			s.logger.Info().Msg("TCP server received LIST command")
			s.DB.List()
		case SET:
			s.logger.Info().Msg("TCP server received SET command")
			if len(cmd.Args) < 2 {
				s.logger.Info().Msg("TCP server SET command needs 2 arguments.")
				s.WriteCommand(cmd.Conn, "Error: SET command needs 2 arguments")
			} else {
				s.DB.Set(cmd.Args)
			}
		default:
			s.WriteCommand(cmd.Conn, "Error: command does not exist")
		}

	}
}
