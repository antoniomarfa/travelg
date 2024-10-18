package ports

import (
	"context"

	"travel/core/models"

	"travel/tools/repository"
)

// UserRepositoy interface
type PermissionRepository interface {
	repository.Repository
}

// SaleService interface
type PermissionService interface {
	Create(ctx context.Context, permisson models.CreateRolesPermissionsReq) (models.CreationResp, error)
	GetAll(ctx context.Context) ([]models.RolesPermissionsResp, error)
	GetByID(ctx context.Context, ID string) (models.RolesPermissionsResp, error)
	Update(ctx context.Context, ID string, permission models.UpdateRolesPermissionsReq) error
	Delete(ctx context.Context, ID string) error
}
