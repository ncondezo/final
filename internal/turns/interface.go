package turns

import (
	"context"

	"github.com/ncondezo/final/internal/domain"
)

type Repository interface {
	Create(ctx context.Context, turn domain.Turn) (domain.Turn, error)
	GetByID(ctx context.Context, id int) (domain.Turn, error)
	GetByPatientID(ctx context.Context, patientId int) ([]domain.Turn, error)
	Update(ctx context.Context, turn domain.Turn, id int) (domain.Turn, error)
	Delete(ctx context.Context, id int) error
}
