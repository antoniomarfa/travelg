package ports

import (
	"context"

	"travel/core/models"

	"travel/tools/repository"
)

// UserRepositoy interface
type ComunasRepository interface {
	repository.Repository
}

// RolesService interface
type ComunasService interface {
	Create(ctx context.Context, comunas models.CreateComunasReq) (models.CreationResp, error)
	GetAll(ctx context.Context, filter map[string]interface{}) ([]models.ComunasResp, error)
	GetByID(ctx context.Context, ID string) (models.ComunasResp, error)
	Update(ctx context.Context, ID string, comunas models.UpdateComunasReq) error
	Delete(ctx context.Context, ID string) error
}
