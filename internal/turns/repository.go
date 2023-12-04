package turns

import (
	"context"
	"database/sql"
	"errors"

	"github.com/go-sql-driver/mysql"
	"github.com/ncondezo/final/internal/domain"
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

// Patch implements Repository.
func (*repository) Patch(ctx context.Context, turn domain.Turn, id int) (domain.Turn, error) {
	panic("unimplemented")
}

func NewRepository(db *sql.DB) Repository {
	return &repository{db: db}
}

// Create is a method that creates a new turn.

func (r *repository) Create(ctx context.Context, turn domain.Turn) (domain.Turn, error) {
	var mysqlError *mysql.MySQLError

	statement, err := r.db.Prepare(QueryInsertTurn)
	if err != nil {
		return domain.Turn{}, ErrPrepareStatement
	}
	defer statement.Close()

	result, err := statement.Exec(
		turn.IdDentist,
		turn.IdDentist,
		turn.Date,
		turn.Description,
	)

	ok := errors.As(err, &mysqlError)
	if ok {
		switch mysqlError.Number {
		case 1062:
			return domain.Turn{}, ErrAlreadyExists
		default:
			return domain.Turn{}, ErrExecStatement
		}
	}

	lastId, err := result.LastInsertId()
	if err != nil {
		return domain.Turn{}, ErrLastInsertedId
	}

	turn.Id = int(lastId)

	return turn, nil
}

// GetByID is a method that returns a turn by ID.
func (r *repository) GetByID(ctx context.Context, id int) (domain.Turn, error) {
	row := r.db.QueryRow(QUeryGetTurnById, id)

	var turn domain.Turn
	err := row.Scan(
		&turn.Id,
		&turn.IdDentist,
		&turn.IdPatient,
		&turn.Date,
		&turn.Description,
	)
	if err == sql.ErrNoRows {
		return domain.Turn{}, ErrNotFound
	}
	if err != nil {
		return domain.Turn{}, ErrExecStatement
	}

	return turn, nil
}

// Update is a method that updates a turn by ID.
func (r *repository) Update(ctx context.Context, turn domain.Turn, id int) (domain.Turn, error) {
	statement, err := r.db.Prepare(QuertyUpdateTurn)
	if err != nil {
		return domain.Turn{}, ErrPrepareStatement
	}
	defer statement.Close()

	_, err = statement.Exec(
		turn.IdDentist,
		turn.IdPatient,
		turn.Date,
		turn.Description,
		id,
	)

	if err != nil {
		return domain.Turn{}, ErrExecStatement
	}

	return turn, nil
}

// Delete is a method that deletes a turn by ID.
func (r *repository) Delete(ctx context.Context, id int) error {
	result, err := r.db.Exec(QueryDeleteTurn, id)
	if err != nil {
		return ErrExecStatement
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected < 1 {
		return ErrNotFound
	}

	return nil
}
