package tcp

import (
	"fmt"
	"strings"
)

type CommandHandlerI interface {
	Handle(cmd Command, s *Server)
}

type ListCommandHandler struct{}

func (h *ListCommandHandler) Handle(cmd Command, s *Server) {
	s.logger.Info().Msg("TCP server received LIST command")
	records := s.DB.List()
	var sb strings.Builder
	for key, value := range records {
		sb.WriteString(fmt.Sprintf("Key: %s, Value: %s\n", key, string(value.Data)))
	}
	response := sb.String()
	if response == "" {
		response = "No records found"
	}
	s.WriteCommand(cmd.Conn, response)
}

type SetCommandHandler struct{}

func (h *SetCommandHandler) Handle(cmd Command, s *Server) {
	s.logger.Info().Msg("TCP server received SET command")
	if len(cmd.Args) != 2 {
		s.WriteCommand(cmd.Conn, "Error: SET command needs 2 arguments")
		return
	}
	val := s.DB.Set(cmd.Args[0], cmd.Args[1]) // TODO: accept expiry

	s.WriteCommand(cmd.Conn, fmt.Sprintf("Set key: %s, value: %s", cmd.Args[0], string(val.Data)))
}

type GetCommandHandler struct{}

func (h *GetCommandHandler) Handle(cmd Command, s *Server) {
	s.logger.Info().Msg("TCP server received GET command")
	if len(cmd.Args) != 1 {
		s.WriteCommand(cmd.Conn, "Error: GET command needs 1 argument")
		return
	}
	value, exists := s.DB.Get(cmd.Args[0])
	if exists {
		s.WriteCommand(cmd.Conn, string(value))
	} else {
		s.WriteCommand(cmd.Conn, "Key not found")
	}
}

type DelCommandHandler struct{}

func (h *DelCommandHandler) Handle(cmd Command, s *Server) {
	s.logger.Info().Msg("TCP server received DEL command")
	if len(cmd.Args) != 1 {
		s.WriteCommand(cmd.Conn, "Error: DEL command needs 1 argument")
		return
	}
	deleted := s.DB.Del(cmd.Args[0])
	if deleted {
		s.WriteCommand(cmd.Conn, "OK")
	} else {
		s.WriteCommand(cmd.Conn, "Key not found")
	}
}

// type ExpireCommandHandler struct{}

// func (h *ExpireCommandHandler) Handle(cmd Command, s *Server) {
// 	s.logger.Info().Msg("TCP server received EXPIRE command")
// 	if len(cmd.Args) < 2 {
// 		s.WriteCommand(cmd.Conn, "Error: EXPIRE command needs 2 arguments")
// 		return
// 	}
// 	duration, err := time.ParseDuration(cmd.Args)
// 	if err != nil {
// 		s.WriteCommand(cmd.Conn, "Error: Invalid duration format")
// 		return
// 	}
// 	updated := s.DB.Expire(cmd.Args[0], duration)
// 	if updated {
// 		s.WriteCommand(cmd.Conn, "OK")
// 	} else {
// 		s.WriteCommand(cmd.Conn, "Key not found")
// 	}
// }
