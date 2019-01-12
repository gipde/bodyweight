package main

import "testing"

const (
	testid   = "TESTID"
	testname = "TESTER"
)

func TestFindItem(t *testing.T) {
	entries := getEntries(testid)
	if entries != nil {
		t.Error("Expected 0 Records. Wrong number of Records:", len(*entries))
	}
}

func TestCreateItem(t *testing.T) {
	for i := 0; i < 10; i++ {
		err := createEntry(testid, testname, TrainingState{Level: 0, Week: 0, Day: 0, Unit:i}, "TestRecord")
		if err != nil {
			t.Error(err)
		}
	}
}
func TestGetItems(t *testing.T) {
	entries := getEntries(testid)
	if len(*entries) != 10 {
		t.Error("Expected 10 Record. Wrong number of Records:", len(*entries))
	}
}

func TestGetLastItem(t *testing.T) {
	entry := getLastUsedEntry(testid)
	if entry.TrainingState.Unit!=9 {
		t.Error("Expected Unit 9, got: ",entry.TrainingState.Unit)
	}
}


func TestDeleteRecords(t *testing.T) {
	err := deleteAllEntries(testid)
	if err != nil {
		t.Error(err)
	}
}
