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
type cursoService struct {
	config     config.Config
	repository ports.CursoRepository
}

// NewURolesService creates a new user service
func NewCursoService(cfg config.Config, repo ports.CursoRepository) ports.CursoService {
	return &cursoService{
		config:     cfg,
		repository: repo,
	}
}

// Create roles
func (p *cursoService) Create(ctx context.Context, curso models.CreateCursoReq) (resp models.CreationResp, err error) {

	now := time.Now().UTC()
	curso.CreatedDate = now
	curso.UpdatedDate = now
	insertedID, err := p.repository.Create(ctx, models.CreateCursoReq(curso))
	if err != nil {
		return
	}

	resp = models.CreationResp{
		InsertedID: insertedID,
	}

	return
}

// GetAll users
func (p *cursoService) GetAll(ctx context.Context, filter map[string]interface{}) (resp []models.CursoResp, err error) {
	// Obtiene los roles desde el repositorio
	result, err := p.repository.Get(ctx, map[string]interface{}{}, nil, nil)
	if err != nil {
		return nil, err
	}

	// Convierte los resultados a la estructura de respuesta
	resp = make([]models.CursoResp, len(result))
	for i, v := range result {
		// Asegúrate de que el tipo sea correcto antes de hacer la conversión
		if curso, ok := v.(models.CursoResp); ok {
			resp[i] = models.CursoResp(curso) // Asumiendo que tienes un constructor o un mapeo en tu struct
		} else {
			return nil, wrappers.NewNonExistentErr(fmt.Errorf("error de tipo: se esperaba entities.Products, pero se obtuvo %T", v))
		}
	}
	return
}

// GetByID user
func (p *cursoService) GetByID(ctx context.Context, ID string) (resp models.CursoResp, err error) {
	curso, err := p.repository.GetByID(ctx, ID)

	if err != nil {
		return resp, fmt.Errorf("pago con ID %s no encontrado", ID)
	}

	if curso == nil {
		// Si no se encuentra el colegio (colegios es nil), devolver un valor en blanco y un error
		return models.CursoResp{}, fmt.Errorf("curso con ID %s no encontrado", ID)
	}

	resp = *curso.(*models.CursoResp)

	return
}

// Update user
func (p *cursoService) Update(ctx context.Context, ID string, curso models.UpdateCursoReq) (err error) {

	dbCurso, err := p.GetByID(ctx, ID)
	if err != nil {
		return
	}
	// Actualizar los campos solo si no son nil
	// Actualizar la fecha de modificación
	dbCurso.UpdatedDate = time.Now().UTC()

	// Llamar al repositorio para actualizar la entidad
	err = p.repository.Update(ctx, ID, models.Curso(dbCurso))

	return err
}

// Delete user
func (p *cursoService) Delete(ctx context.Context, ID string) (err error) {
	err = p.repository.Delete(ctx, ID)
	if errors.Is(err, wrappers.NonExistentErr) {
		err = wrappers.NewNonExistentErr(fmt.Errorf("ID %s not found", ID))
	}

	return err
}
