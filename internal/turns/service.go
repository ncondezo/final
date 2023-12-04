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
	Update(ctx context.Context, dto domain.TurnDTO, id int) (domain.Turn, error)
	Patch(ctx context.Context, dto domain.Turn, id int) (domain.Turn, error)
	Delete(ctx context.Context, id int) error
}

type service struct {
	repository Repository
}

// Patch implements Service.
func (*service) Patch(ctx context.Context, dto domain.Turn, id int) (domain.Turn, error) {
	panic("unimplemented")
}

func NewTurnService(repository Repository) Service {
	return &service{repository: repository}
}

// Create is a method that create a new turn.
func (s *service) Create(ctx context.Context, dto domain.TurnDTO) (domain.Turn, error) {
	turn := domain.Turn{
		IdDentist:   dto.IdDentist,
		IdPatient:   dto.IdPatient,
		Date:        time.Now(),
		Description: dto.Description,
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

// Update is a method that update a turn by ID.
func (s *service) Update(ctx context.Context, dto domain.TurnDTO, id int) (domain.Turn, error) {
	turn, err := s.GetByID(ctx, id)
	if err != nil {
		return domain.Turn{}, err
	}
	turn.IdDentist = dto.IdDentist
	turn.IdPatient = dto.IdPatient
	turn.Description = dto.Description

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
