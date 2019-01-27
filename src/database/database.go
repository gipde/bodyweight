package database

import (
	"time"

	"bodyweight/training"
)

// PK is the Primary Key
type PK struct {
	AlexaID string    `json:"alexa_id"`
	Date    time.Time `json:"date"`
}

// Entry is a database Entry
type Entry struct {
	PK
	UserName      string         `json:"username"`
	TrainingState training.State `json:"training_state"`
}

// DB interface
type DB interface {
	CreateDBIfNotExists() error
	DeleteDB() error

	CreateEntry(entry *Entry) error

	GetLastUsedEntry(alexaID string) (*Entry, error)
	GetEntries(alexaID string) (*[]Entry, error)

	DeleteItem(pk PK) error
	DeleteAllEntries(alexaID string) error
}

// Accessor returns the DB Accessor
func Accessor() *DynamoDB {
	return DynamoDBAccessor()
}
