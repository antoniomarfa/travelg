package ports

import (
	"context"

	"travel/core/models"

	"travel/tools/repository"
)

// UserRepositoy interface
type FmedicaRepository interface {
	repository.Repository
}

// SaleService interface
type FmedicaService interface {
	Create(ctx context.Context, fmedica models.CreateFmedicaReq) (models.CreationResp, error)
	GetAll(ctx context.Context) ([]models.FmedicaResp, error)
	GetByID(ctx context.Context, ID string) (models.FmedicaResp, error)
	Update(ctx context.Context, ID string, fmedica models.UpdateFmedicaReq) error
	Delete(ctx context.Context, ID string) error
}
