package memdb

import (
	"fmt"
	"time"
)

// LIST
func (db DB) List() {
	db.Mutex.Lock()
	defer db.Mutex.Unlock()
	for key, value := range db.Records {
		fmt.Println("key", key, "value", value)
	}

	// return
}

// SET [key] [value]
func (db DB) Set(args []string) {
	db.Mutex.Lock()
	defer db.Mutex.Unlock()
	db.Records[args[0]] = Value{
		Created: time.Now(),
		Expiry:  time.Hour,
		Data:    []byte(args[1]),
	}
}
