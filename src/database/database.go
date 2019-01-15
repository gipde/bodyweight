package database

import (
	"time"

	"bodyweight/training"
)

// Entry is a database Entry
type Entry struct {
	AlexaID       string                 `json:"alexa_id"`
	Date          time.Time              `json:"date"`
	UserName      string                 `json:"username"`
	TrainingState training.State `json:"training_state"`
	Desc          string                 `json:"desc"`
}

// DB interface
type DB interface {
	CreateDBIfNotExists() error
	DeleteDB() error

	CreateEntry(alexaID string, name string, training training.State, desc string) error

	GetLastUsedEntry(alexaID string) (*Entry, error)
	GetEntries(alexaID string) (*[]Entry, error)

	DeleteItem(alexaID string, date time.Time) error
	DeleteAllEntries(alexaID string) error
}
