package ports

import (
	"context"

	"travel/core/models"

	"travel/tools/repository"
)

// UserRepositoy interface
type ProgramRepository interface {
	repository.Repository
}

// SaleService interface
type ProgramService interface {
	Create(ctx context.Context, program models.CreateProgramReq) (models.CreationResp, error)
	GetAll(ctx context.Context) ([]models.ProgramResp, error)
	GetByID(ctx context.Context, ID string) (models.ProgramResp, error)
	Update(ctx context.Context, ID string, program models.UpdateProgramReq) error
	Delete(ctx context.Context, ID string) error
}
