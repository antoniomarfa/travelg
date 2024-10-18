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
type comunasService struct {
	config     config.Config
	repository ports.ComunasRepository
}

// NewURolesService creates a new user service
func NewComunasService(cfg config.Config, repo ports.ComunasRepository) ports.ComunasService {
	return &comunasService{
		config:     cfg,
		repository: repo,
	}
}

// Create roles
func (p *comunasService) Create(ctx context.Context, comunas models.CreateComunasReq) (resp models.CreationResp, err error) {

	now := time.Now().UTC()
	comunas.CreatedDate = now
	comunas.UpdatedDate = now
	insertedID, err := p.repository.Create(ctx, models.CreateComunasReq(comunas))
	if err != nil {
		return
	}

	resp = models.CreationResp{
		InsertedID: insertedID,
	}

	return
}

// GetAll users
func (p *comunasService) GetAll(ctx context.Context, filter map[string]interface{}) (resp []models.ComunasResp, err error) {
	// Obtiene los roles desde el repositorio
	result, err := p.repository.Get(ctx, filter, nil, nil)
	//	result, err := p.repository.Get(ctx, map[string]interface{}{}, nil, nil)
	if err != nil {
		return nil, err
	}

	// Convierte los resultados a la estructura de respuesta
	resp = make([]models.ComunasResp, len(result))
	for i, v := range result {
		// Asegúrate de que el tipo sea correcto antes de hacer la conversión
		if comunas, ok := v.(models.ComunasResp); ok {
			resp[i] = models.ComunasResp(comunas) // Asumiendo que tienes un constructor o un mapeo en tu struct
		} else {
			return nil, wrappers.NewNonExistentErr(fmt.Errorf("error de tipo: se esperaba entities.Products, pero se obtuvo %T", v))
		}
	}
	return
}

// GetByID user
func (p *comunasService) GetByID(ctx context.Context, ID string) (resp models.ComunasResp, err error) {
	comunas, err := p.repository.GetByID(ctx, ID)

	if err != nil {
		return resp, fmt.Errorf("comuna con ID %s no encontrado", ID)
	}

	if comunas == nil {
		// Si no se encuentra el colegio (colegios es nil), devolver un valor en blanco y un error
		return models.ComunasResp{}, fmt.Errorf("comuna con ID %s no encontrado", ID)
	}

	resp = *comunas.(*models.ComunasResp)

	return
}

// Update user
func (p *comunasService) Update(ctx context.Context, ID string, comunas models.UpdateComunasReq) (err error) {

	dbComunas, err := p.GetByID(ctx, ID)
	if err != nil {
		return
	}
	// Actualizar los campos solo si no son nil

	if comunas.Description != nil {
		dbComunas.Description = *comunas.Description
	}

	if comunas.Author != nil {
		dbComunas.Author = *comunas.Author
	}

	// Asegúrate de que Active no sea nil antes de asignarlo
	if comunas.Active != nil {
		dbComunas.Active = *comunas.Active
	}

	// Actualizar la fecha de modificación
	dbComunas.UpdatedDate = time.Now().UTC()

	// Llamar al repositorio para actualizar la entidad
	err = p.repository.Update(ctx, ID, models.Comunas(dbComunas))

	return err
}

// Delete user
func (p *comunasService) Delete(ctx context.Context, ID string) (err error) {
	err = p.repository.Delete(ctx, ID)
	if errors.Is(err, wrappers.NonExistentErr) {
		err = wrappers.NewNonExistentErr(fmt.Errorf("ID %s not found", ID))
	}

	return err
}
