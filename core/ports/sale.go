package ports

import (
	"context"

	"travel/core/models"

	"travel/tools/repository"
)

// UserRepositoy interface
type SaleRepository interface {
	repository.Repository
}

// SaleService interface
type SaleService interface {
	Create(ctx context.Context, sales models.CreateSaleReq) (models.CreationResp, error)
	GetAll(ctx context.Context) ([]models.SaleResp, error)
	GetByID(ctx context.Context, ID string) (models.SaleResp, error)
	Update(ctx context.Context, ID string, sale models.UpdateSaleReq) error
	Delete(ctx context.Context, ID string) error
}
