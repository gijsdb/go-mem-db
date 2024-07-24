package controller

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/chzyer/readline"
	"github.com/rs/zerolog"
)

type CLIController struct {
	logger        zerolog.Logger
	tls_cert_path string
}

func NewCLIController(logger zerolog.Logger, tls_cert_path string) CLIController {
	return CLIController{
		logger:        logger,
		tls_cert_path: tls_cert_path,
	}
}

func (c *CLIController) HandleStartCLI(address string) {
	ssl_config := c.HandleSSLConf()

	conn, err := tls.Dial("tcp", address, ssl_config)
	if err != nil {
		fmt.Printf("Failed to connect to server: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close()

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

		line = strings.TrimSpace(line)
		if line == "exit" {
			break
		}

		_, err = fmt.Fprintf(conn, "%s\n", line)
		if err != nil {
			log.Fatalf("Failed to send command: %v", err)
		}

		response := make([]byte, 4096)
		n, err := conn.Read(response)
		if err != nil {
			log.Fatalf("Failed to read response: %v", err)
		}

		fmt.Println(string(response[:n]))
	}
}

func (c *CLIController) HandleSSLConf() *tls.Config {
	certPool := x509.NewCertPool()
	serverCert, err := os.ReadFile(c.tls_cert_path)
	if err != nil {
		c.logger.Fatal().Err(err).Msg("Cannot find TLS cert path")
	}
	certPool.AppendCertsFromPEM(serverCert)

	return &tls.Config{
		RootCAs: certPool,
	}
}
