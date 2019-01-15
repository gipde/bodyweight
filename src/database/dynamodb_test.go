package database

import (
	"bodyweight/training"
	"log"
	"math/rand"
	"os"
	"strconv"
	"testing"
)

const (
	testid   = "TESTID"
	testname = "TESTER"
)

var db DB

//TODO: test illeagal parameters

func TestMain(m *testing.M) {
	log.Println("starting Tests ... ")

	r := strconv.Itoa(rand.Int())
	tablename := "TESTTABLE_" + r
	log.Println("Creating Table:", tablename)

	db = NewDynamoDB(tablename)

	os.Exit(m.Run())
}

func TestCreateDB(t *testing.T) {
	// DB not exists
	db.CreateDBIfNotExists()
	// DB already exists
	db.CreateDBIfNotExists()
}

func TestFindItem(t *testing.T) {
	entries, _ := db.GetEntries(testid)
	if entries != nil {
		t.Error("Expected 0 Records. Wrong number of Records:", len(*entries))
	}
}

func TestGetLastItem1(t *testing.T) {
	entry, _ := db.GetLastUsedEntry(testid)
	if entry != nil {
		t.Error("Expected nil, got: ", entry)
	}
}

func TestCreateItem(t *testing.T) {
	for i := 0; i < 10; i++ {
		err := db.CreateEntry(testid, testname, training.State{Level: 0, Week: 0, Day: 0, Unit: i}, "TestRecord")
		if err != nil {
			t.Error(err)
		}
	}
}
func TestGetItems2(t *testing.T) {
	entries, _ := db.GetEntries(testid)
	if len(*entries) != 10 {
		t.Error("Expected 10 Record. Wrong number of Records:", len(*entries))
	}
}

func TestGetLastItem2(t *testing.T) {
	entry, _ := db.GetLastUsedEntry(testid)
	if entry.TrainingState.Unit != 9 {
		t.Error("Expected Unit 9, got: ", entry.TrainingState.Unit)
	}
}

func TestDeleteRecords(t *testing.T) {
	err := db.DeleteAllEntries(testid)
	if err != nil {
		t.Error(err)
	}
}

func TestDeleteDB(t *testing.T) {
	err := db.DeleteDB()
	if err != nil {
		t.Error(err)
	}
}
