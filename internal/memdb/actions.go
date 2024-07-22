package memdb

import (
	"time"
)

// LIST
func (db *DB) List() map[string]Value {
	db.Mutex.RLock()
	defer db.Mutex.RUnlock()

	data := make(map[string]Value)
	for key, val := range db.Records {
		data[key] = val
	}

	return data
}

// SET [key] [value] [expiry]
// TODO: Allows setting of multiple key,values
func (db DB) Set(key string, value string, expire time.Duration) Value {
	db.Mutex.Lock()
	defer db.Mutex.Unlock()

	db.Records[key] = Value{
		Created: time.Now(),
		Expiry:  expire,
		Data:    []byte(value),
	}

	return db.Records[key]
}

// GET [key]
func (db *DB) Get(key string) ([]byte, bool) {
	db.Mutex.Lock()
	defer db.Mutex.Unlock()
	value, exists := db.Records[key]
	if !exists {
		return nil, false
	}
	// Check if the key has expired
	if value.Expiry > 0 && time.Since(value.Created) > value.Expiry {
		delete(db.Records, key)
		return nil, false
	}
	return value.Data, true
}

// DEL [key]
func (db *DB) Del(key string) bool {
	db.Mutex.Lock()
	defer db.Mutex.Unlock()
	_, exists := db.Records[key]
	if exists {
		delete(db.Records, key)
		return true
	}
	return false
}

// // EXPIRE [key] [duration]
// func (db *DB) Expire(key string, duration time.Duration) bool {
// 	db.Mutex.Lock()
// 	defer db.Mutex.Unlock()
// 	value, exists := db.Records[key]
// 	if !exists {
// 		return false
// 	}
// 	value.Expiry = duration
// 	db.Records[key] = value
// 	return true
// }
