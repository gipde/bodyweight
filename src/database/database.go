package database

import (
	"time"

	"github.com/gipde/bodyweight/training"
)

// Entry is a database Entry
type Entry struct {
	AlexaID       string                 `json:"alexa_id"`
	Date          time.Time              `json:"date"`
	UserName      string                 `json:"username"`
	TrainingState training.TrainingState `json:"training_state"`
	Desc          string                 `json:"desc"`
}

type DB interface {

	CreateDBIfNotExists() error
	DeleteDB() error

	CreateEntry(alexaID string, name string, training training.TrainingState, desc string) error

	GetLastUsedEntry(alexaID string) (*Entry, error)
	GetEntries(alexaID string) (*[]Entry, error)

	DeleteItem(alexaID string, date time.Time) error
	DeleteAllEntries(alexaID string) error
}
