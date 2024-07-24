package controller

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"

	"github.com/chzyer/readline"
	"github.com/rs/zerolog"
)

type CLIController struct {
	logger zerolog.Logger
}

func NewCLIController(logger zerolog.Logger) CLIController {
	return CLIController{
		logger: logger,
	}
}

func (c *CLIController) HandleStartCLI(address string) {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		fmt.Printf("Failed to connect to server: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close()

	// Create a new readline instance
	rl, err := readline.NewEx(&readline.Config{
		Prompt:          "> ",
		HistoryFile:     "/tmp/readline.tmp",
		AutoComplete:    nil,
		InterruptPrompt: "^C",
		EOFPrompt:       "exit",
	})
	if err != nil {
		c.logger.Fatal().Err(err).Msg("Failed to initialize readline")
	}
	defer rl.Close()

	for {
		// Read input from the user
		line, err := rl.Readline()
		if err == readline.ErrInterrupt {
			if len(line) == 0 {
				break
			} else {
				continue
			}
		} else if err == io.EOF {
			break
		} else if err != nil {
			c.logger.Fatal().Err(err).Msg("Failed to read line")
		}

		// Trim whitespace and check for exit command
		line = strings.TrimSpace(line)
		if line == "exit" {
			break
		}

		// Send the command to the TCP server
		_, err = fmt.Fprintf(conn, "%s\n", line)
		if err != nil {
			log.Fatalf("Failed to send command: %v", err)
		}

		// Read the response from the server
		response := make([]byte, 4096)
		n, err := conn.Read(response)
		if err != nil {
			log.Fatalf("Failed to read response: %v", err)
		}

		// Print the server's response
		fmt.Println(string(response[:n]))
	}

	// reader := bufio.NewReader(os.Stdin)
	// fmt.Println("Connected to server. Type commands:")

	// for {
	// 	fmt.Print("> ")
	// 	input, _ := reader.ReadString('\n')
	// 	input = strings.TrimSpace(input)
	// 	if input == "" {
	// 		continue
	// 	}
	// 	if input == "exit" {
	// 		break
	// 	}

	// 	_, err := conn.Write([]byte(input + "\n"))
	// 	if err != nil {
	// 		fmt.Printf("Failed to send command: %v\n", err)
	// 		break
	// 	}

	// 	response, err := bufio.NewReader(conn).ReadString('\n')

	// 	if err != nil {
	// 		fmt.Printf("Failed to read response: %v\n", err)
	// 		break
	// 	}
	// 	fmt.Println(strings.TrimSpace(response))
	// }
}
