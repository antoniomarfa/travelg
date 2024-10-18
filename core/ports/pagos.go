package ports

import (
	"context"

	"travel/core/models"

	"travel/tools/repository"
)

// UserRepositoy interface
type PagosRepository interface {
	repository.Repository
}

// RolesService interface
type PagosService interface {
	Create(ctx context.Context, pagos models.CreatePagosReq) (models.CreationResp, error)
	GetAll(ctx context.Context) ([]models.PagosResp, error)
	GetByID(ctx context.Context, ID string) (models.PagosResp, error)
	Update(ctx context.Context, ID string, pagos models.UpdatePagosReq) error
	Delete(ctx context.Context, ID string) error
}
