package dentists

import (
	"context"
	"log"

	"github.com/ncondezo/final/internal/domain"
)

type Service interface {
	Create(ctx context.Context, dto domain.DentistDTO) (domain.Dentist, error)
	GetByID(ctx context.Context, id int) (domain.Dentist, error)
	Update(ctx context.Context, dto domain.DentistDTO, id int) (domain.Dentist, error)
	Patch(ctx context.Context, dto domain.DentistDTO, id int) (domain.Dentist, error)
	Delete(ctx context.Context, id int) error
}

type service struct {
	repository Repository
}

func NewDentistService(repository Repository) Service {
	return &service{repository: repository}
}

// Create is a method that create a new dentist.
func (s *service) Create(ctx context.Context, dto domain.DentistDTO) (domain.Dentist, error) {
	dentist := domain.Dentist{
		Name:         dto.Name,
		LastName:     dto.LastName,
		Registration: dto.Registration,
	}
	dentist, err := s.repository.Create(ctx, dentist)
	if err != nil {
		log.Println("[DentistService][Create] error creating dentist", err)
		return domain.Dentist{}, err
	}
	return dentist, nil
}

// GetByID is a method that return a dentist by ID.
func (s *service) GetByID(ctx context.Context, id int) (domain.Dentist, error) {
	dentist, err := s.repository.GetByID(ctx, id)
	if err != nil {
		log.Println("[DentistService][GetById] error getting dentist", err)
		return domain.Dentist{}, err
	}
	return dentist, nil
}

// Update is a method that update a dentist by ID.
func (s *service) Update(ctx context.Context, dto domain.DentistDTO, id int) (domain.Dentist, error) {
	dentist := domain.Dentist{
		Name:         dto.Name,
		LastName:     dto.LastName,
		Registration: dto.Registration,
	}
	dentist, err := s.repository.Update(ctx, dentist, id)
	if err != nil {
		log.Println("[DentistService][Update] error updating dentist", err)
		return domain.Dentist{}, err
	}
	return dentist, nil
}

// Patch is a method that update a dentist by ID.
func (s *service) Patch(ctx context.Context, dto domain.DentistDTO, id int) (domain.Dentist, error) {
	dentist := domain.Dentist{
		Name:         dto.Name,
		LastName:     dto.LastName,
		Registration: dto.Registration,
	}
	dentist, err := s.repository.Patch(ctx, dentist, id)
	if err != nil {
		log.Println("[DentistService][Patch] error patching dentist", err)
		return domain.Dentist{}, err
	}
	return dentist, nil
}

// Delete is a method that delete a dentist by ID.
func (s *service) Delete(ctx context.Context, id int) error {
	err := s.repository.Delete(ctx, id)
	if err != nil {
		log.Println("[DentistService][Delete] error deleting dentist", err)
		return err
	}
	return nil
}
