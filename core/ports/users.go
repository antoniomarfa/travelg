package ports

import (
	"context"

	"travel/core/models"

	"travel/tools/repository"
)

// UserRepositoy interface
type UsersRepository interface {
	repository.Repository
}

// SaleService interface
type UsersService interface {
	Create(ctx context.Context, users models.CreateUsersReq) (models.CreationResp, error)
	GetAll(ctx context.Context) ([]models.UsersResp, error)
	GetByID(ctx context.Context, ID string) (models.UsersResp, error)
	Update(ctx context.Context, ID string, users models.UpdateUsersReq) error
	Delete(ctx context.Context, ID string) error
}
