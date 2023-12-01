package patients

import (
	"context"
	"log"

	"github.com/ncondezo/final/internal/domain"
)

type Service interface {
	Create(ctx context.Context, pacient domain.Patient) (domain.Patient, error)
	GetByID(ctx context.Context, id int) (domain.Patient, error)
	Update(ctx context.Context, patient domain.Patient, id int) (domain.Patient, error)
	Patch(ctx context.Context, patient domain.Patient, id int) (domain.Patient, error)
	Delete(ctx context.Context, id int) error
}

type service struct {
	repository Repository
}

func NewPatientService(repository Repository) Service {
	return &service{repository: repository}
}

// Create is a method that create a new patient.
func (s *service) Create(ctx context.Context, patient domain.Patient) (domain.Patient, error) {
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
func (s *service) Update(ctx context.Context, patient domain.Patient, id int) (domain.Patient, error) {
	patient, err := s.repository.Update(ctx, patient, id)
	if err != nil {
		log.Println("[PatientsService][Update] error updating patient", err)
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

// Patch is a method that update a patient by ID.
func (s *service) Patch(ctx context.Context, patient domain.Patient, id int) (domain.Patient, error) {
	patientStored, err := s.repository.GetByID(ctx, id)
	if err != nil {
		log.Println("[PatientsService][Patch] error getting patient", err)
		return domain.Patient{}, err
	}

	patientNew, err := s.validate(patientStored, patient)
	if err != nil {
		log.Println("[PatientsService][Patch] error validating patient", err)
		return domain.Patient{}, err
	}

	patient, err = s.repository.Patch(ctx, patientNew, id)
	if err != nil {
		log.Println("[PatientsService][Patch] error patching patient", err)
		return domain.Patient{}, err
	}

	return patient, nil
}

// Validate is a method that validate the fields from body request.
func (s *service) validate(patientStored, dto domain.Patient) (domain.Patient, error) {
	if dto.Name != "" {
		patientStored.Name = dto.Name
	}
	if dto.Lastname != "" {
		patientStored.Lastname = dto.Lastname
	}
	if dto.Address != "" {
		patientStored.Address = dto.Address
	}
	if dto.Dni != "" {
		patientStored.Dni = dto.Dni
	}
	return patientStored, nil
}
