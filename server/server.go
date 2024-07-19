package server

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"

	"github.com/gijsdb/go-mem-db/memdb"
)

type Command struct {
	Value string
	Args  []string
	Conn  net.Conn
}

type Server struct {
	Addr     string
	Listener net.Listener
	DB       memdb.DB
	Commands chan Command
}

func NewServer(addr string) Server {
	return Server{
		Addr:     addr,
		DB:       memdb.NewDB(),
		Commands: make(chan Command),
	}
}

func (s *Server) Start() {
	listener, err := net.Listen("tcp", s.Addr)
	if err != nil {
		panic(err)
	}
	defer listener.Close()
	s.Listener = listener

	s.HandleConnections()
}

func (s *Server) HandleConnections() {
	for {
		conn, err := s.Listener.Accept()
		if err != nil {
			log.Printf("unable to accept connection: %v", err)
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
			log.Printf("unable to read from connection: %v", err)
			break
		}
		cmd := strings.Trim(input, "\r\n")
		args := strings.Split(cmd, " ")

		s.Commands <- Command{
			Value: args[0], Args: args[1:], Conn: conn,
		}
	}
}

func (s *Server) HandleCommand() {
	for cmd := range s.Commands {
		switch cmd.Value {
		case LIST:
			fmt.Println("LIST")
			s.DB.List()
		case SET:
			fmt.Println("SET")
			s.DB.Set(cmd.Args)
		}
	}
}
