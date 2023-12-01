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

func (s *service) Create(ctx context.Context, dentist domain.Dentist) (domain.Dentist, error) {
	dentist, err := s.repository.Create(ctx, dentist)
	if err != nil {
		log.Println("[DentistService][Create] error creating dentist", err)
		return domain.Dentist{}, err
	}
	return dentist, nil
}

func (s *service) GetByID(ctx context.Context, id int) (domain.Dentist, error) {
	dentist, err := s.repository.GetByID(ctx, id)
	if err != nil {
		log.Println("[DentistService][GetById] error getting dentist by id", err)
		return domain.Dentist{}, err
	}
	return dentist, nil
}

func (s *service) Update(ctx context.Context, dentist domain.Dentist, id int) (domain.Dentist, error) {
	dentist, err := s.repository.Update(ctx, dentist, id)
	if err != nil {
		log.Println("[DentistService][Update] error updating dentist", err)
		return domain.Dentist{}, err
	}
	return dentist, nil
}

func (s *service) Delete(ctx context.Context, id int) error {
	err := s.repository.Delete(ctx, id)
	if err != nil {
		log.Println("[DentistService][Delete] error deleting dentist", err)
		return err
	}
	return nil
}

func (s *service) Patch(ctx context.Context, dentist domain.Dentist, id int) (domain.Dentist, error) {
	repositoryDentist, err := s.repository.GetByID(ctx, id)
	if err != nil {
		log.Println("[DentistService][Patch] error getting dentist by ID", err)
		return domain.Dentist{}, err
	}

	dentistPatch, err := s.validate(repositoryDentist, dentist)
	if err != nil {
		log.Println("[DentistService][Patch] error validating dentist", err)
		return domain.Dentist{}, err
	}

	dentist, err = s.repository.Patch(ctx, dentistPatch, id)
	if err != nil {
		log.Println("[DentistService][Patch] error patching dentist by ID", err)
		return domain.Dentist{}, err
	}
	return dentist, nil
}

func (s *service) validate(repositoryDentist, dentist domain.Dentist) (domain.Dentist, error) {

	if dentist.Name != "" {
		repositoryDentist.Name = dentist.Name
	}

	if dentist.LastName != "" {
		repositoryDentist.LastName = dentist.LastName
	}

	if dentist.Registration != "" {
		repositoryDentist.Registration = dentist.Registration
	}
	return repositoryDentist, nil
}
