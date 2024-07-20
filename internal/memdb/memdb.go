package memdb

import (
	"sync"
	"time"

	"github.com/rs/zerolog"
)

// Maybe an ID to create multiple in future? Manage with web UI.
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

func NewDB(logger zerolog.Logger) DB {
	logger.Info().Msg("memdb database created")
	return DB{
		logger:  logger,
		Mutex:   &sync.RWMutex{},
		Records: make(map[string]Value),
	}
}
