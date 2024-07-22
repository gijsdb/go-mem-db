package memdb

import (
	"testing"
	"time"

	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
)

var test_data = map[string]Value{
	"1": {Created: time.Now(), Expiry: time.Hour, Data: []byte("value1")},
	"2": {Created: time.Now(), Expiry: time.Hour, Data: []byte("value2")},
}

func Setup() DB {
	db := NewDB(zerolog.Logger{})
	db.Records = test_data
	return *db
}

func TestList(t *testing.T) {
	db := Setup()

	res := db.List()

	assert.Equal(t, test_data, res)
}

func TestSet(t *testing.T) {
	db := Setup()

	res := db.Set("3", "value3")

	assert.Equal(t, db.Records["3"], res)
}

func TestGet(t *testing.T) {
	db := Setup()

	res, found := db.Get("1")

	assert.True(t, found)
	assert.Equal(t, test_data["1"].Data, res)
}

func TestDel(t *testing.T) {
	db := Setup()

	res := db.Del("1")

	assert.True(t, res)
	assert.NotContains(t, db.Records, test_data["1"])
}
