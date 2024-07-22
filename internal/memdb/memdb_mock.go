package memdb

import "time"

type MockDB struct {
	Records map[string]Value
}

func NewMockDB() *MockDB {
	return &MockDB{Records: make(map[string]Value)}
}

func (db *MockDB) List() map[string]Value {
	data := make(map[string]Value)
	for key, val := range db.Records {
		data[key] = val
	}
	return data
}

func (db *MockDB) Set(key string, value string, expire time.Duration) Value {
	val := Value{
		Created: time.Now(),
		Expiry:  expire,
		Data:    []byte(value),
	}
	db.Records[key] = val
	return val
}

func (db *MockDB) Get(key string) ([]byte, bool) {
	val, exists := db.Records[key]
	if !exists {
		return nil, false
	}
	return val.Data, true
}

func (db *MockDB) Del(key string) bool {
	_, exists := db.Records[key]
	if exists {
		delete(db.Records, key)
	}
	return exists
}

// func (db *MockDB) Expire(key string, duration time.Duration) bool {
// 	val, exists := db.Records[key]
// 	if exists {
// 		val.Expiry = duration
// 		db.Records[key] = val
// 	}
// 	return exists
// }
