package memdb

import (
	"fmt"
	"sync"
	"time"
)

// Maybe an ID to create multiple in future? Manage with web UI.
type DB struct {
	Mutex   *sync.RWMutex
	Records map[string]Value
}

type Value struct {
	Created time.Time
	Expiry  time.Duration
	Data    []byte
}

func NewDB() DB {
	fmt.Println("Creating new database")
	return DB{
		Mutex:   &sync.RWMutex{},
		Records: make(map[string]Value),
	}
}
