package patients

import (
	"context"
	"log"
	"time"

	"github.com/ncondezo/final/internal/domain"
)

type Service interface {
	Create(ctx context.Context, pacient domain.Patient) (domain.Patient, error)
	GetByID(ctx context.Context, id int) (domain.Patient, error)
	Update(ctx context.Context, patient domain.Patient, id int) (domain.Patient, error)
	Delete(ctx context.Context, id int) error
	Patch(ctx context.Context, patient domain.Patient, id int) (domain.Patient, error)
}

type service struct {
	repository Repository
}

func NewServicePatient(repository Repository) Service {
	return &service{repository: repository}
}

// Create is a method that creates a new patient.
func (s *service) Create(ctx context.Context, patient domain.Patient) (domain.Patient, error) {
	patient, err := s.repository.Create(ctx, patient)
	if err != nil {
		log.Println("[PatientsService][Create] error creating patient", err)
		return domain.Patient{}, err
	}

	return patient, nil
}

// GetByID is a method that returns a pacient by ID.
func (s *service) GetByID(ctx context.Context, id int) (domain.Patient, error) {
	patient, err := s.repository.GetByID(ctx, id)
	if err != nil {
		log.Println("[PatientService][GetByID] error getting patient by ID", err)
		return domain.Patient{}, err
	}

	return patient, nil
}

// Update is a method that updates a patient by ID.
func (s *service) Update(ctx context.Context, patient domain.Patient, id int) (domain.Patient, error) {
	patient, err := s.repository.Update(ctx, patient, id)
	if err != nil {
		log.Println("[PatientsService][Update] error updating patient by ID", err)
		return domain.Patient{}, err
	}

	return patient, nil
}

// Delete is a method that deletes a patient by ID.
func (s *service) Delete(ctx context.Context, id int) error {
	err := s.repository.Delete(ctx, id)
	if err != nil {
		log.Println("[PatientsService][Delete] error deleting patient by ID", err)
		return err
	}

	return nil
}

// Patch is a method that updates a patient by ID.
func (s *service) Patch(ctx context.Context, patient domain.Patient, id int) (domain.Patient, error) {
	patientStore, err := s.repository.GetByID(ctx, id)
	if err != nil {
		log.Println("[PatientsService][Patch] error getting patient by ID", err)
		return domain.Patient{}, err
	}

	patientPatch, err := s.validatePatch(patientStore, patient)
	if err != nil {
		log.Println("[PatientsService][Patch] error validating patient", err)
		return domain.Patient{}, err
	}

	patient, err = s.repository.Patch(ctx, patientPatch, id)
	if err != nil {
		log.Println("[PatientsService][Patch] error patching patient by ID", err)
		return domain.Patient{}, err
	}

	return patient, nil
}

// validatePatch is a method that validates the fields to be updated.
func (s *service) validatePatch(patientStore, patient domain.Patient) (domain.Patient, error) {

	if patient.Name != "" {
		patientStore.Name = patient.Name
	}

	if patient.Lastname != "" {
		patientStore.Lastname = patient.Lastname
	}

	if patient.Address != "" {
		patientStore.Address = patient.Address
	}

	if patient.Dni != "" {
		patientStore.Dni = patient.Dni
	}

	if !patient.DateUp.Equal(time.Time{}) {
		patientStore.DateUp = patient.DateUp
	}

	return patientStore, nil

}
