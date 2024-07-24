package controller

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"os"

	"github.com/gijsdb/go-mem-db/internal/memdb"
	"github.com/gijsdb/go-mem-db/internal/tcp"
	"github.com/rs/zerolog"
)

type TCPServerController struct {
	logger   zerolog.Logger
	address  string
	port     string
	tls_conf *tls.Config
	db       *memdb.DB
}

func NewTCPServerController(logger zerolog.Logger, address string, port string, ssl_cert_path string, ssl_key_path string, db *memdb.DB) TCPServerController {
	c := TCPServerController{
		logger:  logger,
		address: address,
		port:    port,
		db:      db,
	}
	c.tls_conf = c.HandleSSLConf(ssl_cert_path, ssl_key_path)

	return c
}

func (c *TCPServerController) HandleStartTCPServer() {
	s := tcp.NewServer(fmt.Sprintf("%s:%s", c.address, c.port), c.db, c.logger)
	s.Start(c.tls_conf)
}

func (s *TCPServerController) HandleSSLConf(cert_path, key_path string) *tls.Config {
	cert, err := tls.LoadX509KeyPair(cert_path, key_path)
	if err != nil {
		s.logger.Fatal().Err(err).Msg("server cannot load SSL key pair")
	}

	certPool := x509.NewCertPool()
	serverCert, err := os.ReadFile("cert.pem")
	if err != nil {
		s.logger.Fatal().Err(err).Msg("server cannot read SSL cert")
	}
	certPool.AppendCertsFromPEM(serverCert)

	return &tls.Config{
		Certificates: []tls.Certificate{cert},
		ClientCAs:    certPool,
	}
}
