package services

import (
	"context"
	"errors"
	"fmt"

	"travel/config"
	"travel/core/models"
	"travel/core/ports"
	"travel/tools/wrappers"
)

// rolesService adapter of an user service
type voucherService struct {
	config     config.Config
	repository ports.PagosRepository
}

// NewURolesService creates a new user service
func NewVoucherService(cfg config.Config, repo ports.VoucherRepository) ports.VoucherService {
	return &voucherService{
		config:     cfg,
		repository: repo,
	}
}

// Create roles
func (p *voucherService) Create(ctx context.Context, voucher models.CreateVoucherReq) (resp models.CreationResp, err error) {

	insertedID, err := p.repository.Create(ctx, models.CreateVoucherReq(voucher))
	if err != nil {
		return
	}

	resp = models.CreationResp{
		InsertedID: insertedID,
	}

	return
}

// GetAll users
func (p *voucherService) GetAll(ctx context.Context, filter map[string]interface{}) (resp []models.VoucherResp, err error) {
	// Obtiene los roles desde el repositorio
	result, err := p.repository.Get(ctx, map[string]interface{}{}, nil, nil)
	if err != nil {
		return nil, err
	}

	// Convierte los resultados a la estructura de respuesta
	resp = make([]models.VoucherResp, len(result))
	for i, v := range result {
		// Asegúrate de que el tipo sea correcto antes de hacer la conversión
		if voucher, ok := v.(models.VoucherResp); ok {
			resp[i] = models.VoucherResp(voucher) // Asumiendo que tienes un constructor o un mapeo en tu struct
		} else {
			return nil, wrappers.NewNonExistentErr(fmt.Errorf("error de tipo: se esperaba entities.Products, pero se obtuvo %T", v))
		}
	}
	return
}

// GetByID user
func (p *voucherService) GetByID(ctx context.Context, ID string) (resp models.VoucherResp, err error) {
	voucher, err := p.repository.GetByID(ctx, ID)

	if err != nil {
		return resp, fmt.Errorf("voucher con ID %s no encontrado", ID)
	}

	if voucher == nil {
		// Si no se encuentra el colegio (colegios es nil), devolver un valor en blanco y un error
		return models.VoucherResp{}, fmt.Errorf("voucher con ID %s no encontrado", ID)
	}

	resp = *voucher.(*models.VoucherResp)

	return
}

// Update user
func (p *voucherService) Update(ctx context.Context, ID string, voucher models.UpdateVoucherReq) (err error) {

	dbVoucher, err := p.GetByID(ctx, ID)
	if err != nil {
		return
	}
	// Actualizar los campos solo si no son nil
	// Actualizar la fecha de modificación
	dbVoucher.Used = voucher.Used

	// Llamar al repositorio para actualizar la entidad
	err = p.repository.Update(ctx, ID, models.Voucher(dbVoucher))

	return err
}

// Delete user
func (p *voucherService) Delete(ctx context.Context, ID string) (err error) {
	err = p.repository.Delete(ctx, ID)
	if errors.Is(err, wrappers.NonExistentErr) {
		err = wrappers.NewNonExistentErr(fmt.Errorf("ID %s not found", ID))
	}

	return err
}
