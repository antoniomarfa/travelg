package ports

import (
	"context"

	"travel/core/models"

	"travel/tools/repository"
)

// IngresoRepositoy interface
type IngresoRepository interface {
	repository.Repository
}

// IngresoService interface
type IngresoService interface {
	Create(ctx context.Context, ingreso models.CreateIngresoReq) (models.CreationResp, error)
	GetAll(ctx context.Context) ([]models.IngresoResp, error)
	GetByID(ctx context.Context, ID string) (models.IngresoResp, error)
	Update(ctx context.Context, ID string, ingreso models.UpdateIngresoReq) error
	Delete(ctx context.Context, ID string) error
}
