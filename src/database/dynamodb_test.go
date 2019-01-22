package database

import (
	"bodyweight/tools"
	"bodyweight/training"
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"testing"
	"time"
)

const (
	testid   = "TESTID"
	testname = "TESTER"
)

var db *DynamoDB

func TestMain(m *testing.M) {
	log.Println("starting Tests ... ")
	tools.SetupTestDB()
	db = Accessor()

	os.Exit(m.Run())
}

func TestCreateDB(t *testing.T) {
	// DB already exists, because it is created in TestMain
	err := db.CreateDBIfNotExists()
	assert.NoError(t, err)
}

func TestFindItem(t *testing.T) {
	entries, _ := db.GetEntries(testid)
	assert.Nil(t, entries)
}

func TestGetLastItem1(t *testing.T) {
	entry, err := db.GetLastUsedEntry(testid)
	assert.NoError(t, err)
	assert.Nil(t, entry)
}

func TestCreateItem(t *testing.T) {
	for i := 0; i < 10; i++ {
		entry := Entry{
			AlexaID:  testid,
			UserName: testname,
			Date:     time.Now(),
			TrainingState: training.State{
				Level: 0,
				Week:  0,
				Day:   0,
				Unit:  i,
			},
			Desc: "TestRecord",
		}
		err := db.CreateEntry(&entry)
		assert.NoError(t, err)
	}
}
func TestGetItems2(t *testing.T) {
	entries, _ := db.GetEntries(testid)
	assert.Equal(t, 10, len(*entries), "Expected 10 Record. Wrong number of Records.")
}

func TestGetLastItem2(t *testing.T) {
	entry, _ := db.GetLastUsedEntry(testid)
	assert.Equal(t, 9, entry.TrainingState.Unit)
}

func TestDeleteRecords(t *testing.T) {
	err := db.DeleteAllEntries(testid)
	assert.NoError(t, err)
}

func TestUnknown(t *testing.T) {
	e := db.DeleteItem("unknown", time.Now())
	log.Println("error:", e)
}

func TestDeleteDB(t *testing.T) {
	err := db.DeleteDB()
	assert.NoError(t, err)
}
