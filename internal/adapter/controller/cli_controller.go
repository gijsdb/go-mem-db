package controller

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"

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

func (c *CLIController) HandleStartCLI() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: cli <address>")
		os.Exit(1)
	}
	address := os.Args[1]

	conn, err := net.Dial("tcp", address)
	if err != nil {
		fmt.Printf("Failed to connect to server: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close()

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Connected to server. Type commands:")

	for {
		fmt.Print("> ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		if input == "" {
			continue
		}
		if input == "exit" {
			break
		}

		_, err := conn.Write([]byte(input + "\n"))
		if err != nil {
			fmt.Printf("Failed to send command: %v\n", err)
			break
		}

		response, err := bufio.NewReader(conn).ReadString('\n')

		if err != nil {
			fmt.Printf("Failed to read response: %v\n", err)
			break
		}
		fmt.Println(strings.TrimSpace(response))
	}
}
