package tcp

import (
	"time"

	"github.com/gijsdb/go-mem-db/internal/memdb"
	"github.com/rs/zerolog"
)

var test_data = map[string]memdb.Value{
	"1": {Created: time.Now(), Expiry: time.Hour, Data: []byte("value1")},
	"2": {Created: time.Now(), Expiry: time.Hour, Data: []byte("value2")},
}

func SetUp() *Server {
	logger := zerolog.Logger{}
	db := memdb.NewDB(zerolog.Logger{})
	db.Records = test_data

	s := NewServer("localhost:1111", db, logger)

	return &s
}
