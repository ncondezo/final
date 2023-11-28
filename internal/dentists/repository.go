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
)

type repositorymysql struct {
	db *sql.DB
}

// NewMemoryRepository ....
func NewMySqlRepository(db *sql.DB) Repository {
	return &repositorymysql{db: db}
}

// Create a new dentist
func (r *repositorymysql) Create(ctx context.Context, dentist domain.Dentist) (domain.Dentist, error) {
	statement, err := r.db.Prepare(QueryInsertDentist)
	if err != nil {
		return dentist, ErrPrepareStatement
	}

	defer statement.Close()

	result, err := statement.Exec(
		dentist.LastName,
		dentist.Name,
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

// Get a dentist by id
func (r *repositorymysql) GetByID(ctx context.Context, id int) (domain.Dentist, error) {
	row := r.db.QueryRow(QueryGetDentistById, id)

	var dentist domain.Dentist
	err := row.Scan(
		&dentist.Id,
		&dentist.LastName,
		&dentist.Name,
		&dentist.Registration,
	)

	if err != nil {
		return domain.Dentist{}, err
	}

	return dentist, nil
}

// Update dentist info
func (r *repositorymysql) Update(
	ctx context.Context,
	dentist domain.Dentist,
	id int) (domain.Dentist, error) {
	statement, err := r.db.Prepare(QueryUpdateDentist)
	if err != nil {
		return domain.Dentist{}, err
	}

	defer statement.Close()

	result, err := statement.Exec(
		dentist.LastName,
		dentist.Name,
		dentist.Registration,
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

func (r *repositorymysql) Patch(
	ctx context.Context,
	dentist domain.Dentist,
	id int) (domain.Dentist, error) {
	statement, err := r.db.Prepare(QueryUpdateDentist)
	if err != nil {
		return domain.Dentist{}, err
	}

	defer statement.Close()

	result, err := statement.Exec(
		dentist.LastName,
		dentist.Name,
		dentist.Registration,
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

func (r *repositorymysql) Delete(ctx context.Context, id int) error {
	result, err := r.db.Exec(QueryDeleteDentist, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected < 1 {
		return errors.New("e")
	}

	return nil
}
