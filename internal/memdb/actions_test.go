package memdb

import (
	"testing"
	"time"

	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
)

// TODO: test edge cases

var test_data = map[string]Value{
	"1":       {Created: time.Now(), Expiry: time.Hour, Data: []byte("value1")},
	"2":       {Created: time.Now(), Expiry: time.Hour, Data: []byte("value2")},
	"expired": {Created: time.Now().Add(-10 * time.Minute), Expiry: time.Minute * 1, Data: []byte("value2")},
}

func Setup() DB {
	db := NewDB(zerolog.Logger{})
	db.Records = test_data
	return *db
}

func TestList(t *testing.T) {
	db := Setup()

	data := db.List()

	assert.Equal(t, test_data, data)
}

func Test_Set_Sets_Value(t *testing.T) {
	db := Setup()

	data := db.Set("3", "value3", time.Minute*10)

	assert.Equal(t, db.Records["3"], data)
}

func Test_Get_Gets_Value(t *testing.T) {
	db := Setup()

	data, found := db.Get("1")

	assert.True(t, found)
	assert.Equal(t, test_data["1"].Data, data)
}

func Test_Get_Does_Not_Get_Value_Not_exist(t *testing.T) {
	db := Setup()

	_, found := db.Get("no_exist")
	assert.False(t, found)
}

func Test_Get_Does_Not_Get_Expired_Value(t *testing.T) {
	db := Setup()

	data, found := db.Get("expired")
	assert.False(t, found)
	assert.Nil(t, data)
}

func Test_Del_Deletes_Value(t *testing.T) {
	db := Setup()

	success := db.Del("1")

	assert.True(t, success)
	assert.NotContains(t, db.Records, test_data["1"])
}
