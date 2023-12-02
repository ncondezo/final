package turns

import (
	"database/sql"
	"errors"
)

var (
	ErrPrepareStatement = errors.New("error prepare statement")
	ErrExecStatement    = errors.New("error exec statement")
	ErrLastInsertedId   = errors.New("error last inserted id")
	ErrNotFound         = errors.New("error not found turn")
	ErrAlreadyExists    = errors.New("error turn already exists")
)

type repository struct {
	db *sql.DB
}

// Create is a method that creates a new turn.
