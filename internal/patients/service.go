package patients

import (
	"context"
	"log"
	"time"

	"github.com/ncondezo/final/internal/domain"
)

type Service interface {
	Create(ctx context.Context, dto domain.PatientDTO) (domain.Patient, error)
	GetByID(ctx context.Context, id int) (domain.Patient, error)
	Update(ctx context.Context, dto domain.PatientDTO, id int) (domain.Patient, error)
	Patch(ctx context.Context, dto domain.PatientDniDTO, id int) (domain.Patient, error)
	Delete(ctx context.Context, id int) error
}

type service struct {
	repository Repository
}

func NewPatientService(repository Repository) Service {
	return &service{repository: repository}
}

// Create is a method that create a new patient.
func (s *service) Create(ctx context.Context, dto domain.PatientDTO) (domain.Patient, error) {
	patient := domain.Patient{
		Name:     dto.Name,
		Lastname: dto.Lastname,
		Address:  dto.Address,
		Dni:      dto.Dni,
		DateUp:   time.Now(),
	}
	patient, err := s.repository.Create(ctx, patient)
	if err != nil {
		log.Println("[PatientsService][Create] error creating patient", err)
		return domain.Patient{}, err
	}
	return patient, nil
}

// GetByID is a method that return a patient by ID.
func (s *service) GetByID(ctx context.Context, id int) (domain.Patient, error) {
	patient, err := s.repository.GetByID(ctx, id)
	if err != nil {
		log.Println("[PatientService][GetByID] error getting patient", err)
		return domain.Patient{}, err
	}
	return patient, nil
}

// Update is a method that update a patient by ID.
func (s *service) Update(ctx context.Context, dto domain.PatientDTO, id int) (domain.Patient, error) {
	patient, err := s.GetByID(ctx, id)
	if err != nil {
		return domain.Patient{}, err
	}
	patient.Name = dto.Name
	patient.Lastname = dto.Lastname
	patient.Address = dto.Address
	patient.Dni = dto.Dni
	patient, err = s.repository.Update(ctx, patient, id)
	if err != nil {
		log.Println("[PatientsService][Update] error updating patient", err)
		return domain.Patient{}, err
	}
	return patient, nil
}

// Patch is a method that update a patient dni by ID.
func (s *service) Patch(ctx context.Context, dto domain.PatientDniDTO, id int) (domain.Patient, error) {
	patient, err := s.GetByID(ctx, id)
	if err != nil {
		return domain.Patient{}, err
	}
	patient.Dni = dto.Dni
	patient, err = s.repository.Patch(ctx, patient, id)
	if err != nil {
		log.Println("[PatientsService][Patch] error patching patient", err)
		return domain.Patient{}, err
	}
	return patient, nil
}

// Delete is a method that delete a patient by ID.
func (s *service) Delete(ctx context.Context, id int) error {
	err := s.repository.Delete(ctx, id)
	if err != nil {
		log.Println("[PatientsService][Delete] error deleting patient", err)
		return err
	}
	return nil
}
