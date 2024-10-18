package ports

import (
	"context"

	"travel/core/models"

	"travel/tools/repository"
)

// UserRepositoy interface
type ColegiosRepository interface {
	repository.Repository
}

// RolesService interface
type ColegiosService interface {
	Create(ctx context.Context, colegios models.CreateColegiosReq) (models.CreationResp, error)
	GetAll(ctx context.Context) ([]models.ColegiosResp, error)
	GetByID(ctx context.Context, ID string) (models.ColegiosResp, error)
	Update(ctx context.Context, ID string, company models.UpdateColegiosReq) error
	Delete(ctx context.Context, ID string) error
}
