package tcp

import (
	"testing"
	"time"

	"github.com/gijsdb/go-mem-db/internal/memdb"
	"github.com/rs/zerolog"
)

func SetUp() Server {
	logger := zerolog.New(nil)
	mockDB := memdb.NewMockDB()
	mockDB.Records["key1"] = memdb.Value{Data: []byte("value1")}
	mockDB.Records["key2"] = memdb.Value{Data: []byte("value2")}
	return NewServer("127.0.0.1:4242", mockDB, logger)
}

func Test_List_Command_Handler_Lists_Results(t *testing.T) {
	server := SetUp()

	go server.HandleCommand()

	mockConn := NewMockConn("LIST\n")
	go server.ReadCommand(mockConn)

	time.Sleep(100 * time.Millisecond)

	expectedOutput := "Key: key1, Value: value1\nKey: key2, Value: value2\n"
	if output := mockConn.GetWrittenData(); output != expectedOutput {
		t.Errorf("Expected output %q, got %q", expectedOutput, output)
	}
}

func Test_List_Command_Handler_Errors_With_Args(t *testing.T) {
	server := SetUp()

	go server.HandleCommand()

	mockConn := NewMockConn("LIST arg\n")
	go server.ReadCommand(mockConn)

	time.Sleep(100 * time.Millisecond)

	expectedOutput := "Error: LIST command does not take arguments"
	if output := mockConn.GetWrittenData(); output != expectedOutput {
		t.Errorf("Expected output %q, got %q", expectedOutput, output)
	}
}

func Test_Set_Command_Handler_Sets_Value(t *testing.T) {
	server := SetUp()

	go server.HandleCommand()

	mockConn := NewMockConn("SET 1 value1 10\n")
	go server.ReadCommand(mockConn)

	time.Sleep(100 * time.Millisecond)

	expectedOutput := "Set key: 1, value: value1"
	if output := mockConn.GetWrittenData(); output != expectedOutput {
		t.Errorf("Expected output %q, got %q", expectedOutput, output)
	}
}

func Test_Set_Command_Handler_Only_Accepts_3_Args(t *testing.T) {
	server := SetUp()

	go server.HandleCommand()

	mockConn := NewMockConn("SET 1 value1 10 arg\n")
	go server.ReadCommand(mockConn)

	time.Sleep(100 * time.Millisecond)

	expectedOutput := "Error: SET command needs 3 arguments"
	if output := mockConn.GetWrittenData(); output != expectedOutput {
		t.Errorf("Expected output %q, got %q", expectedOutput, output)
	}

	mockConn = NewMockConn("SET 1 value1\n")
	go server.ReadCommand(mockConn)

	time.Sleep(100 * time.Millisecond)

	if output := mockConn.GetWrittenData(); output != expectedOutput {
		t.Errorf("Expected output %q, got %q", expectedOutput, output)
	}
}

func Test_Get_Command_Handler_Gets_Value(t *testing.T) {
	server := SetUp()

	go server.HandleCommand()

	mockConn := NewMockConn("GET key1\n")
	go server.ReadCommand(mockConn)

	time.Sleep(100 * time.Millisecond)

	expectedOutput := "value1"
	if output := mockConn.GetWrittenData(); output != expectedOutput {
		t.Errorf("Expected output %q, got %q", expectedOutput, output)
	}
}

func Test_Get_Command_Handler_Only_Accepts_1_Arg(t *testing.T) {
	server := SetUp()

	go server.HandleCommand()
	expectedOutput := "Error: GET command needs 1 argument"

	mockConn := NewMockConn("GET key1 key2\n")
	go server.ReadCommand(mockConn)

	time.Sleep(100 * time.Millisecond)

	if output := mockConn.GetWrittenData(); output != expectedOutput {
		t.Errorf("Expected output %q, got %q", expectedOutput, output)
	}

	mockConn = NewMockConn("GET\n")
	go server.ReadCommand(mockConn)

	time.Sleep(100 * time.Millisecond)

	if output := mockConn.GetWrittenData(); output != expectedOutput {
		t.Errorf("Expected output %q, got %q", expectedOutput, output)
	}
}

func Test_Del_Command_Handler_Deletes_Value(t *testing.T) {
	server := SetUp()

	go server.HandleCommand()

	mockConn := NewMockConn("DEL key1\n")
	go server.ReadCommand(mockConn)

	time.Sleep(100 * time.Millisecond)

	expectedOutput := "OK"
	if output := mockConn.GetWrittenData(); output != expectedOutput {
		t.Errorf("Expected output %q, got %q", expectedOutput, output)
	}
}

func Test_Del_Command_Handler_Only_Accepts_1_Arg(t *testing.T) {
	server := SetUp()

	go server.HandleCommand()
	expectedOutput := "Error: DEL command needs 1 argument"

	mockConn := NewMockConn("DEL key1 key1\n")
	go server.ReadCommand(mockConn)

	time.Sleep(100 * time.Millisecond)

	if output := mockConn.GetWrittenData(); output != expectedOutput {
		t.Errorf("Expected output %q, got %q", expectedOutput, output)
	}

	mockConn = NewMockConn("DEL\n")
	go server.ReadCommand(mockConn)

	time.Sleep(100 * time.Millisecond)

	if output := mockConn.GetWrittenData(); output != expectedOutput {
		t.Errorf("Expected output %q, got %q", expectedOutput, output)
	}
}
