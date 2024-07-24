package tcp

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

const (
	LIST  = "LIST"
	GET   = "GET"
	SET   = "SET"
	DEL   = "DEL"
	PATCH = "SET"
)

type Command struct {
	Value string
	Args  []string
	Conn  net.Conn
}

var CommandHandlerMap = map[string]CommandHandlerI{
	"LIST": &ListCommandHandler{},
	"SET":  &SetCommandHandler{},
	"GET":  &GetCommandHandler{},
	"DEL":  &DelCommandHandler{},
	// "EXPIRE": &ExpireCommandHandler{},
}

func (s *Server) ReadCommand(conn net.Conn) {
	defer conn.Close()
	for {
		input, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			s.logger.Info().Msg(fmt.Sprintf("no message received on connection... closing: %v", err))
			break
		}
		cmd := strings.Trim(input, "\r\n")
		args := strings.Split(cmd, " ")

		s.commands <- Command{
			Value: args[0], Args: args[1:], Conn: conn,
		}
	}
}

func (s *Server) WriteCommand(conn net.Conn, data string) {
	if _, err := conn.Write([]byte(data)); err != nil {
		s.logger.Error().Msg(fmt.Sprintf("TCP server failed to write : %v", err))
		conn.Close()
	}
}

func (s *Server) HandleCommand() {
	for cmd := range s.commands {
		handler, exists := CommandHandlerMap[cmd.Value]
		if exists {
			handler.Handle(cmd, s)
		} else {
			s.WriteCommand(cmd.Conn, "Error: Unknown command")
		}
	}
}
