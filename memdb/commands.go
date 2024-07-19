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
}

// SET [key] [value]
func (db DB) Set(args []string) {
	db.Mutex.Lock()
	defer db.Mutex.Unlock()
	if len(args) > 2 {
		fmt.Println("too many args")
	} else {
		db.Records[args[0]] = Value{
			Created: time.Now(),
			Expiry:  time.Hour,
			Data:    []byte(args[1]),
		}
	}
}
