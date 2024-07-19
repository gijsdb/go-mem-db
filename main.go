package main

import (
	"github.com/gijsdb/go-mem-db/server"
)

func main() {
	s := server.NewServer("localhost:4242")
	s.Start()
}
