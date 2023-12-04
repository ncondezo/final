package turns

import (
	"context"
	"database/sql"
	"errors"

	"github.com/ncondezo/final/internal/dentists"
	"github.com/ncondezo/final/internal/domain"
	"github.com/ncondezo/final/internal/patients"
)

var (
	ErrPrepareStatement = errors.New("error prepare statement")
	ErrExecStatement    = errors.New("error exec statement")
	ErrLastInsertedId   = errors.New("error last inserted id")
	ErrNotFound         = errors.New("error not found turn")
)

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{db: db}
}

// Create is a method that creates a new turn.
func (r *repository) Create(ctx context.Context, turn domain.Turn) (domain.Turn, error) {

	_, err := patients.NewRepository(r.db).GetByID(ctx, turn.Patient.Id)
	if err != nil {
		return domain.Turn{}, err
	}

	_, err = dentists.NewRepository(r.db).GetByID(ctx, turn.Dentist.Id)
	if err != nil {
		return domain.Turn{}, err
	}

	statement, err := r.db.Prepare(QueryInsertTurn)
	if err != nil {
		return domain.Turn{}, ErrPrepareStatement
	}
	defer statement.Close()

	result, err := statement.Exec(
		turn.Date,
		turn.Description,
		turn.Patient.Id,
		turn.Dentist.Id,
	)
	if err != nil {
		return domain.Turn{}, ErrExecStatement
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
	var ignored int
	var ignored2 int
	row := r.db.QueryRow(QueryGetTurnById, id)

	var turn domain.Turn
	err := row.Scan(
		&turn.Id,
		&turn.Date,
		&turn.Description,
		&ignored,
		&ignored2,
		&turn.Patient.Id,
		&turn.Patient.Name,
		&turn.Patient.Lastname,
		&turn.Patient.Address,
		&turn.Patient.Dni,
		&turn.Patient.DateUp,
		&turn.Dentist.Id,
		&turn.Dentist.Name,
		&turn.Dentist.LastName,
		&turn.Dentist.Registration,
	)
	if errors.Is(err, sql.ErrNoRows) {
		return domain.Turn{}, ErrNotFound
	}
	if err != nil {
		return domain.Turn{}, ErrExecStatement
	}

	return turn, nil
}

// GetByPatientID is a method that returns a list of turns by PatientID.
func (r *repository) GetByPatientID(ctx context.Context, patientId int) ([]domain.Turn, error) {
	turns := make([]domain.Turn, 0)

	_, err := patients.NewRepository(r.db).GetByID(ctx, patientId)
	if err != nil {
		return []domain.Turn{}, err
	}

	founds, err := r.db.Query(QueryGetTurnByPatient, patientId)
	defer founds.Close()
	if err != nil {
		return []domain.Turn{}, ErrExecStatement
	}

	for founds.Next() {
		var turn domain.Turn
		var ignored int
		err := founds.Scan(
			&turn.Id,
			&turn.Date,
			&turn.Description,
			&ignored,
			&ignored,
			&turn.Patient.Id,
			&turn.Patient.Name,
			&turn.Patient.Lastname,
			&turn.Patient.Address,
			&turn.Patient.Dni,
			&turn.Patient.DateUp,
			&turn.Dentist.Id,
			&turn.Dentist.Name,
			&turn.Dentist.LastName,
			&turn.Dentist.Registration,
		)
		if err != nil {
			return []domain.Turn{}, ErrExecStatement
		}
		turns = append(turns, turn)
	}

	return turns, nil
}

// Update is a method that updates a turn by ID.
func (r *repository) Update(ctx context.Context, turn domain.Turn, id int) (domain.Turn, error) {
	_, err := dentists.NewRepository(r.db).GetByID(ctx, turn.Dentist.Id)
	if err != nil {
		return domain.Turn{}, err
	}

	statement, err := r.db.Prepare(QueryUpdateTurn)
	if err != nil {
		return domain.Turn{}, ErrPrepareStatement
	}
	defer statement.Close()

	_, err = statement.Exec(
		turn.Date,
		turn.Description,
		turn.Dentist.Id,
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
