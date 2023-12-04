package turns

import (
	"context"
	"log"
	"time"

	"github.com/ncondezo/final/internal/domain"
)

type Service interface {
	Create(ctx context.Context, dto domain.TurnDTO) (domain.Turn, error)
	GetByID(ctx context.Context, id int) (domain.Turn, error)
	GetByPatientID(ctx context.Context, patientId int) ([]domain.Turn, error)
	Update(ctx context.Context, dto domain.TurnDTO, id int) (domain.Turn, error)
	Delete(ctx context.Context, id int) error
}

type service struct {
	repository Repository
}

func NewTurnService(repository Repository) Service {
	return &service{repository: repository}
}

// Create is a method that create a new turn.
func (s *service) Create(ctx context.Context, dto domain.TurnDTO) (domain.Turn, error) {
	turn := domain.Turn{
		Date:        time.Now(),
		Description: dto.Description,
		Patient: domain.Patient{
			Id: dto.IdPatient,
		},
		Dentist: domain.Dentist{
			Id: dto.IdDentist,
		},
	}
	turn, err := s.repository.Create(ctx, turn)
	if err != nil {
		log.Println("[TurnsService][Create] error creating turn", err)
		return domain.Turn{}, err
	}
	return turn, nil
}

// GetByID is a method that return a turn by ID.
func (s *service) GetByID(ctx context.Context, id int) (domain.Turn, error) {
	turn, err := s.repository.GetByID(ctx, id)
	if err != nil {
		log.Println("[TurnsService][GetByID] error getting turn", err)
		return domain.Turn{}, err
	}
	return turn, nil
}

// GetByPatientID is a method that return a turn by ID.
func (s *service) GetByPatientID(ctx context.Context, patientId int) ([]domain.Turn, error) {
	turns, err := s.repository.GetByPatientID(ctx, patientId)
	if err != nil {
		log.Println("[TurnsService][GetByID] error getting turns by patient", err)
		return []domain.Turn{}, err
	}
	return turns, nil
}

// Update is a method that update a turn by ID.
func (s *service) Update(ctx context.Context, dto domain.TurnDTO, id int) (domain.Turn, error) {
	turn, err := s.GetByID(ctx, id)
	if err != nil {
		return domain.Turn{}, err
	}
	turn.Date = dto.Date
	turn.Description = dto.Description
	turn.Dentist = domain.Dentist{
		Id: dto.IdDentist,
	}

	turn, err = s.repository.Update(ctx, turn, id)
	if err != nil {
		log.Println("[TurnsService][Update] error updating turn", err)
		return domain.Turn{}, err
	}
	return turn, nil
}

// Delete is a method that delete a turn by ID.
func (s *service) Delete(ctx context.Context, id int) error {
	err := s.repository.Delete(ctx, id)
	if err != nil {
		log.Println("[TurnsService][Delete] error deleting turns", err)
		return err
	}
	return nil
}
