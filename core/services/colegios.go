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
type colegiosService struct {
	config     config.Config
	repository ports.CompanyRepository
}

// NewURolesService creates a new user service
func NewColegiosService(cfg config.Config, repo ports.ColegiosRepository) ports.ColegiosService {
	return &colegiosService{
		config:     cfg,
		repository: repo,
	}
}

// Create roles
func (p *colegiosService) Create(ctx context.Context, colegios models.CreateColegiosReq) (resp models.CreationResp, err error) {

	insertedID, err := p.repository.Create(ctx, models.CreateColegiosReq(colegios))
	if err != nil {
		return
	}

	resp = models.CreationResp{
		InsertedID: insertedID,
	}

	return
}

// GetAll users
func (p *colegiosService) GetAll(ctx context.Context) (resp []models.ColegiosResp, err error) {
	// Obtiene los roles desde el repositorio
	result, err := p.repository.Get(ctx, map[string]interface{}{}, nil, nil)
	if err != nil {
		return nil, err
	}

	// Convierte los resultados a la estructura de respuesta
	resp = make([]models.ColegiosResp, len(result))
	for i, v := range result {
		// Asegúrate de que el tipo sea correcto antes de hacer la conversión
		if colegios, ok := v.(models.ColegiosResp); ok {
			resp[i] = models.ColegiosResp(colegios) // Asumiendo que tienes un constructor o un mapeo en tu struct
		} else {
			return nil, wrappers.NewNonExistentErr(fmt.Errorf("error de tipo: se esperaba entities.Colegios, pero se obtuvo %T", v))
		}
	}

	return
}

// GetByID user
func (p *colegiosService) GetByID(ctx context.Context, ID string) (resp models.ColegiosResp, err error) {
	colegios, err := p.repository.GetByID(ctx, ID)

	if err != nil {
		return resp, fmt.Errorf("colegio con ID %s no encontrado", ID)
	}

	if colegios == nil {
		// Si no se encuentra el colegio (colegios es nil), devolver un valor en blanco y un error
		return models.ColegiosResp{}, fmt.Errorf("colegio con ID %s no encontrado", ID)
	}

	//	resp = models.ColegiosResp(*colegios.(*models.ColegiosResp))
	resp = *colegios.(*models.ColegiosResp)
	return resp, nil
	// return
}

// Update user
func (p *colegiosService) Update(ctx context.Context, ID string, colegios models.UpdateColegiosReq) (err error) {
	dbColegios, err := p.GetByID(ctx, ID)
	if err != nil {
		return
	}

	if colegios.Codigo != nil {
		dbColegios.Codigo = *colegios.Codigo
	}

	if colegios.Nombre != nil {
		dbColegios.Nombre = *colegios.Nombre
	}

	if colegios.Direccion != nil {
		dbColegios.Direccion = *colegios.Direccion
	}

	if colegios.Comuna != nil {
		dbColegios.Comuna = *colegios.Comuna
	}

	if colegios.Latitud != nil {
		dbColegios.Latitud = *colegios.Latitud
	}

	if colegios.Longitud != nil {
		dbColegios.Longitud = *colegios.Longitud
	}

	if colegios.RegionId != nil {
		dbColegios.RegionId = *colegios.RegionId
	}

	if colegios.CompanyId != nil {
		dbColegios.ComunaId = *colegios.ComunaId
	}

	if colegios.CompanyId != nil {
		dbColegios.CompanyId = *colegios.CompanyId
	}

	err = p.repository.Update(ctx, ID, models.Colegios(dbColegios))
	return err

}

// Delete user
func (p *colegiosService) Delete(ctx context.Context, ID string) (err error) {
	err = p.repository.Delete(ctx, ID)
	if errors.Is(err, wrappers.NonExistentErr) {
		err = wrappers.NewNonExistentErr(fmt.Errorf("ID %s not found", ID))
	}

	return err
}
