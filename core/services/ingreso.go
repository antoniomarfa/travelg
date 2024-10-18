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

// ingresoService adapter of an user service
type ingresoService struct {
	config     config.Config
	repository ports.IngresoRepository
}

// NewURolesService creates a new user service
func NewIngresoService(cfg config.Config, repo ports.IngresoRepository) ports.IngresoService {
	return &ingresoService{
		config:     cfg,
		repository: repo,
	}
}

// Create roles
func (p *ingresoService) Create(ctx context.Context, ingreso models.CreateIngresoReq) (resp models.CreationResp, err error) {

	now := time.Now().UTC()
	ingreso.CreatedDate = now
	ingreso.UpdatedDate = now
	insertedID, err := p.repository.Create(ctx, models.CreateIngresoReq(ingreso))
	if err != nil {
		return
	}

	resp = models.CreationResp{
		InsertedID: insertedID,
	}

	return
}

// GetAll users
func (p *ingresoService) GetAll(ctx context.Context) (resp []models.IngresoResp, err error) {
	// Obtiene los roles desde el repositorio
	result, err := p.repository.Get(ctx, map[string]interface{}{}, nil, nil)
	if err != nil {
		return nil, err
	}

	// Convierte los resultados a la estructura de respuesta
	resp = make([]models.IngresoResp, len(result))
	for i, v := range result {
		// Asegúrate de que el tipo sea correcto antes de hacer la conversión
		if ingreso, ok := v.(models.IngresoResp); ok {
			resp[i] = models.IngresoResp(ingreso) // Asumiendo que tienes un constructor o un mapeo en tu struct
		} else {
			return nil, wrappers.NewNonExistentErr(fmt.Errorf("error de tipo: se esperaba entities.Products, pero se obtuvo %T", v))
		}
	}
	return
}

// GetByID user
func (p *ingresoService) GetByID(ctx context.Context, ID string) (resp models.IngresoResp, err error) {
	ingreso, err := p.repository.GetByID(ctx, ID)

	if err != nil {
		return resp, fmt.Errorf("ingreso con ID %s no encontrado", ID)
	}

	if ingreso == nil {
		// Si no se encuentra el colegio (colegios es nil), devolver un valor en blanco y un error
		return models.IngresoResp{}, fmt.Errorf("ingreso con ID %s no encontrado", ID)
	}

	resp = *ingreso.(*models.IngresoResp)

	return
}

// Update user
func (p *ingresoService) Update(ctx context.Context, ID string, ingreso models.UpdateIngresoReq) (err error) {

	dbIngreso, err := p.GetByID(ctx, ID)
	if err != nil {
		return
	}
	// Actualizar los campos solo si no son nil
	// Actualizar la fecha de modificación
	dbIngreso.UpdatedDate = time.Now().UTC()

	// Llamar al repositorio para actualizar la entidad
	err = p.repository.Update(ctx, ID, models.Ingreso(dbIngreso))

	return err
}

// Delete user
func (p *ingresoService) Delete(ctx context.Context, ID string) (err error) {
	err = p.repository.Delete(ctx, ID)
	if errors.Is(err, wrappers.NonExistentErr) {
		err = wrappers.NewNonExistentErr(fmt.Errorf("ID %s not found", ID))
	}

	return err
}
