package dentists

import (
	"context"
	"log"

	"github.com/ncondezo/final/internal/domain"
)

type Service interface {
	Create(ctx context.Context, dentist domain.Dentist) (domain.Dentist, error)
	GetByID(ctx context.Context, id int) (domain.Dentist, error)
	Update(ctx context.Context, dentist domain.Dentist, id int) (domain.Dentist, error)
	Patch(ctx context.Context, dentist domain.Dentist, id int) (domain.Dentist, error)
	Delete(ctx context.Context, id int) error
}

type service struct {
	repository Repository
}

func NewDentistService(repository Repository) Service {
	return &service{repository: repository}
}

// Create is a method that create a new dentist.
func (s *service) Create(ctx context.Context, dentist domain.Dentist) (domain.Dentist, error) {
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
func (s *service) Update(ctx context.Context, dentist domain.Dentist, id int) (domain.Dentist, error) {
	dentist, err := s.repository.Update(ctx, dentist, id)
	if err != nil {
		log.Println("[DentistService][Update] error updating dentist", err)
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

// Patch is a method that update a dentist by ID.
func (s *service) Patch(ctx context.Context, dentist domain.Dentist, id int) (domain.Dentist, error) {
	dentistStored, err := s.repository.GetByID(ctx, id)
	if err != nil {
		log.Println("[DentistService][Patch] error getting dentist", err)
		return domain.Dentist{}, err
	}

	dentistNew, err := s.validate(dentistStored, dentist)
	if err != nil {
		log.Println("[DentistService][Patch] error validating dentist", err)
		return domain.Dentist{}, err
	}

	dentist, err = s.repository.Patch(ctx, dentistNew, id)
	if err != nil {
		log.Println("[DentistService][Patch] error patching dentist", err)
		return domain.Dentist{}, err
	}
	return dentist, nil
}

// Validate is a method that validate the fields from body request.
func (s *service) validate(dentistStored, dto domain.Dentist) (domain.Dentist, error) {
	if dto.Name != "" {
		dentistStored.Name = dto.Name
	}
	if dto.LastName != "" {
		dentistStored.LastName = dto.LastName
	}
	if dto.Registration != "" {
		dentistStored.Registration = dto.Registration
	}
	return dentistStored, nil
}
