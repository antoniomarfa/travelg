package services

import (
	"context"
	"errors"
	"fmt"
	"time"

	"travel/config"
	"travel/core/models"
	"travel/core/ports"
	"travel/tools/wrappers"
)

// rolesService adapter of an user service
type pagosService struct {
	config     config.Config
	repository ports.PagosRepository
}

// NewURolesService creates a new user service
func NewPagosService(cfg config.Config, repo ports.PagosRepository) ports.PagosService {
	return &pagosService{
		config:     cfg,
		repository: repo,
	}
}

// Create roles
func (p *pagosService) Create(ctx context.Context, pagos models.CreatePagosReq) (resp models.CreationResp, err error) {

	now := time.Now().UTC()
	pagos.CreatedDate = now
	pagos.UpdatedDate = now
	insertedID, err := p.repository.Create(ctx, models.CreatePagosReq(pagos))
	if err != nil {
		return
	}

	resp = models.CreationResp{
		InsertedID: insertedID,
	}

	return
}

// GetAll users
func (p *pagosService) GetAll(ctx context.Context, filter map[string]interface{}) (resp []models.PagosResp, err error) {
	// Obtiene los roles desde el repositorio
	result, err := p.repository.Get(ctx, map[string]interface{}{}, nil, nil)
	if err != nil {
		return nil, err
	}

	// Convierte los resultados a la estructura de respuesta
	resp = make([]models.PagosResp, len(result))
	for i, v := range result {
		// Asegúrate de que el tipo sea correcto antes de hacer la conversión
		if pagos, ok := v.(models.PagosResp); ok {
			resp[i] = models.PagosResp(pagos) // Asumiendo que tienes un constructor o un mapeo en tu struct
		} else {
			return nil, wrappers.NewNonExistentErr(fmt.Errorf("error de tipo: se esperaba entities.Products, pero se obtuvo %T", v))
		}
	}
	return
}

// GetByID user
func (p *pagosService) GetByID(ctx context.Context, ID string) (resp models.PagosResp, err error) {
	pagos, err := p.repository.GetByID(ctx, ID)

	if err != nil {
		return resp, fmt.Errorf("pago con ID %s no encontrado", ID)
	}

	if pagos == nil {
		// Si no se encuentra el colegio (colegios es nil), devolver un valor en blanco y un error
		return models.PagosResp{}, fmt.Errorf("pago con ID %s no encontrado", ID)
	}

	resp = *pagos.(*models.PagosResp)

	return
}

// Update user
func (p *pagosService) Update(ctx context.Context, ID string, pagos models.UpdatePagosReq) (err error) {

	dbPagos, err := p.GetByID(ctx, ID)
	if err != nil {
		return
	}
	// Actualizar los campos solo si no son nil
	// Actualizar la fecha de modificación
	dbPagos.UpdatedDate = time.Now().UTC()

	// Llamar al repositorio para actualizar la entidad
	err = p.repository.Update(ctx, ID, models.Pagos(dbPagos))

	return err
}

// Delete user
func (p *pagosService) Delete(ctx context.Context, ID string) (err error) {
	err = p.repository.Delete(ctx, ID)
	if errors.Is(err, wrappers.NonExistentErr) {
		err = wrappers.NewNonExistentErr(fmt.Errorf("ID %s not found", ID))
	}

	return err
}
