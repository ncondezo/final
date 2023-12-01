package dentists

import (
	"context"
	"database/sql"
	"errors"

	"github.com/ncondezo/final/internal/domain"
)

var (
	ErrPrepareStatement = errors.New("error prepare statement")
	ErrExecStatement    = errors.New("error exec statement")
	ErrLastInsertedId   = errors.New("error last inserted id")
	ErrNotFound         = errors.New("error not found dentist")
)

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{db: db}
}

// Create is a method that creates a new dentist.
func (r *repository) Create(ctx context.Context, dentist domain.Dentist) (domain.Dentist, error) {
	statement, err := r.db.Prepare(QueryInsertDentist)
	if err != nil {
		return domain.Dentist{}, ErrPrepareStatement
	}
	defer statement.Close()

	result, err := statement.Exec(
		dentist.Id,
		dentist.Name,
		dentist.LastName,
		dentist.Registration,
	)

	if err != nil {
		return domain.Dentist{}, ErrExecStatement
	}

	lastId, err := result.LastInsertId()
	if err != nil {
		return domain.Dentist{}, ErrLastInsertedId
	}

	dentist.Id = int(lastId)

	return dentist, nil
}

// GetByID is a method that returns a dentist by ID.
func (r *repository) GetByID(ctx context.Context, id int) (domain.Dentist, error) {
	row := r.db.QueryRow(QueryGetDentistById, id)

	var dentist domain.Dentist
	err := row.Scan(
		&dentist.Id,
		&dentist.Name,
		&dentist.LastName,
		&dentist.Registration,
	)

	if err != nil {
		return domain.Dentist{}, err
	}

	return dentist, nil
}

// Update is a method that updates a dentist by ID.
func (r *repository) Update(ctx context.Context, dentist domain.Dentist, id int) (domain.Dentist, error) {
	statement, err := r.db.Prepare(QueryUpdateDentist)
	if err != nil {
		return domain.Dentist{}, err
	}

	defer statement.Close()

	result, err := statement.Exec(
		dentist.Name,
		dentist.LastName,
		dentist.Registration,
		id,
	)

	if err != nil {
		return domain.Dentist{}, err
	}

	_, err = result.RowsAffected()
	if err != nil {
		return domain.Dentist{}, err
	}

	dentist.Id = id

	return dentist, nil
}

// Patch is a method that updates a dentist by ID.
func (r *repository) Patch(ctx context.Context, dentist domain.Dentist, id int) (domain.Dentist, error) {
	statement, err := r.db.Prepare(QueryUpdateDentist)
	if err != nil {
		return domain.Dentist{}, err
	}

	defer statement.Close()

	result, err := statement.Exec(
		dentist.Name,
		dentist.LastName,
		dentist.Registration,
		id,
	)

	if err != nil {
		return domain.Dentist{}, err
	}

	_, err = result.RowsAffected()
	if err != nil {
		return domain.Dentist{}, err
	}

	return dentist, nil
}

// Delete is a method that deletes a patient by ID.
func (r *repository) Delete(ctx context.Context, id int) error {
	result, err := r.db.Exec(QueryDeleteDentist, id)
	if err != nil {
		return err
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
