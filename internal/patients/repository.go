package patients

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
	ErrNotFound         = errors.New("error not found patient")
)

type repositorymysql struct {
	db *sql.DB
}

func NewMyRepository(db *sql.DB) Repository {
	return &repositorymysql{db: db}
}

// Create is a method that creates a new pacient.
func (r *repositorymysql) Create(ctx context.Context, patient domain.Patient) (domain.Patient, error) {
	statement, err := r.db.Prepare(QueryInsertPatient)
	if err != nil {
		return domain.Patient{}, ErrPrepareStatement
	}

	defer statement.Close()

	result, err := statement.Exec(
		patient.Id,
		patient.Name,
		patient.Lastname,
		patient.Address,
		patient.Dni,
		patient.DateUp,
	)

	if err != nil {
		return domain.Patient{}, ErrExecStatement
	}

	lastId, err := result.LastInsertId()
	if err != nil {
		return domain.Patient{}, ErrLastInsertedId
	}

	patient.Id = int(lastId)

	return patient, nil

}

// GetByID is a method that returns a pacient by ID.
func (r *repositorymysql) GetByID(ctx context.Context, id int) (domain.Patient, error) {
	row := r.db.QueryRow(QueryGetPatientById, id)

	var patient domain.Patient
	err := row.Scan(
		&patient.Id,
		&patient.Name,
		&patient.Lastname,
		&patient.Address,
		&patient.Dni,
		&patient.DateUp,
	)

	if err != nil {
		return domain.Patient{}, err
	}

	return patient, nil
}

// Update is a method that updates a pacient by ID.
func (r *repositorymysql) Update(
	ctx context.Context,
	patient domain.Patient,
	id int) (domain.Patient, error) {
	statement, err := r.db.Prepare(QueryUpdatePatient)
	if err != nil {
		return domain.Patient{}, err
	}

	defer statement.Close()

	result, err := statement.Exec(
		patient.Name,
		patient.Lastname,
		patient.Address,
		patient.Dni,
		patient.DateUp,
	)

	if err != nil {
		return domain.Patient{}, err
	}

	_, err = result.RowsAffected()
	if err != nil {
		return domain.Patient{}, err
	}

	patient.Id = id

	return patient, nil

}

// Delete is a method that deletes a pacient by ID.
func (r *repositorymysql) Delete(ctx context.Context, id int) error {
	result, err := r.db.Exec(QueryDeletePatient, id)
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

// Patch is a method that updates a pacient by ID.
func (r *repositorymysql) Patch(
	ctx context.Context,
	patient domain.Patient,
	id int) (domain.Patient, error) {
	statement, err := r.db.Prepare(QueryUpdatePatient)
	if err != nil {
		return domain.Patient{}, err
	}

	defer statement.Close()

	result, err := statement.Exec(
		patient.Name,
		patient.Lastname,
		patient.Address,
		patient.Dni,
		patient.DateUp,
	)

	if err != nil {
		return domain.Patient{}, err
	}

	_, err = result.RowsAffected()
	if err != nil {
		return domain.Patient{}, err
	}

	return patient, nil
}
