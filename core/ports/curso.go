package ports

import (
	"context"

	"travel/core/models"

	"travel/tools/repository"
)

// UserRepositoy interface
type CursoRepository interface {
	repository.Repository
}

// SaleService interface
type CursoService interface {
	Create(ctx context.Context, curso models.CreateCursoReq) (models.CreationResp, error)
	GetAll(ctx context.Context) ([]models.CursoResp, error)
	GetByID(ctx context.Context, ID string) (models.CursoResp, error)
	Update(ctx context.Context, ID string, curso models.UpdateCursoReq) error
	Delete(ctx context.Context, ID string) error
}
