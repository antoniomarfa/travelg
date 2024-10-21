package ports

import (
	"context"

	"travel/core/models"

	"travel/tools/repository"
)

// UserRepositoy interface
type VoucherRepository interface {
	repository.Repository
}

// VoucherService interface
type VoucherService interface {
	Create(ctx context.Context, voucher models.CreateVoucherReq) (models.CreationResp, error)
	GetAll(ctx context.Context, filter map[string]interface{}) ([]models.VoucherResp, error)
	GetByID(ctx context.Context, ID string) (models.VoucherResp, error)
	Update(ctx context.Context, ID string, voucher models.UpdateVoucherReq) error
	Delete(ctx context.Context, ID string) error
}
