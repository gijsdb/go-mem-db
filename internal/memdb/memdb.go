package memdb

import (
	"sync"
	"time"

	"github.com/rs/zerolog"
)

type DBI interface {
	List() map[string]Value
	Set(key string, value string) Value // TODO: add expiry
	Get(key string) ([]byte, bool)
	Del(key string) bool
	// Expire(key string, duration time.Duration) bool
}

type DB struct {
	logger  zerolog.Logger
	Mutex   *sync.RWMutex
	Records map[string]Value
}

type Value struct {
	Created time.Time
	Expiry  time.Duration
	Data    []byte
}

func NewDB(logger zerolog.Logger) *DB {
	logger.Info().Msg("memdb database created")
	return &DB{
		logger:  logger,
		Mutex:   &sync.RWMutex{},
		Records: make(map[string]Value),
	}
}
