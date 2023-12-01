package dentists

import (
	"context"

	"github.com/ncondezo/final/internal/domain"
)

type Repository interface {
	Create(ctx context.Context, dentist domain.Dentist) (domain.Dentist, error)
	GetByID(ctx context.Context, id int) (domain.Dentist, error)
	Update(ctx context.Context, dentist domain.Dentist, id int) (domain.Dentist, error)
	Patch(ctx context.Context, dentist domain.Dentist, id int) (domain.Dentist, error)
	Delete(ctx context.Context, id int) error
}
