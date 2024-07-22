package tcp

import (
	"testing"
	"time"

	"github.com/gijsdb/go-mem-db/internal/memdb"
	"github.com/rs/zerolog"
)

func TestListCommandHandler(t *testing.T) {
	logger := zerolog.New(nil)
	mockDB := memdb.NewMockDB()
	mockDB.Records["key1"] = memdb.Value{Data: []byte("value1")}
	mockDB.Records["key2"] = memdb.Value{Data: []byte("value2")}

	server := NewServer("127.0.0.1:4242", mockDB, logger)
	go server.HandleCommand()

	mockConn := NewMockConn("LIST\n")
	go server.ReadCommand(mockConn)

	time.Sleep(100 * time.Millisecond)

	expectedOutput := "Key: key1, Value: value1\nKey: key2, Value: value2\n\n"
	if output := mockConn.GetWrittenData(); output != expectedOutput {
		t.Errorf("Expected output %q, got %q", expectedOutput, output)
	}
}
