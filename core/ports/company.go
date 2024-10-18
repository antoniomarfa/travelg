package ports

import (
	"context"

	"travel/core/models"

	"travel/tools/repository"
)

// UserRepositoy interface
type CompanyRepository interface {
	repository.Repository
}

// RolesService interface
type CompanyService interface {
	Create(ctx context.Context, company models.CreateCompanyReq) (models.CreationResp, error)
	GetAll(ctx context.Context) ([]models.CompanyResp, error)
	GetByID(ctx context.Context, ID string) (models.CompanyResp, error)
	Update(ctx context.Context, ID string, company models.UpdateCompanyReq) error
	Delete(ctx context.Context, ID string) error
}
