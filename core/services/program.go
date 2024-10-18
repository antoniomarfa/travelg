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
type programService struct {
	config     config.Config
	repository ports.ProgramRepository
}

// NewURolesService creates a new user service
func NewProgramService(cfg config.Config, repo ports.ProgramRepository) ports.ProgramService {
	return &programService{
		config:     cfg,
		repository: repo,
	}
}

// Create roles
func (p *programService) Create(ctx context.Context, program models.CreateProgramReq) (resp models.CreationResp, err error) {

	insertedID, err := p.repository.Create(ctx, models.CreateProgramReq(program))
	if err != nil {
		return
	}

	resp = models.CreationResp{
		InsertedID: insertedID,
	}

	return
}

// GetAll users
func (p *programService) GetAll(ctx context.Context) (resp []models.ProgramResp, err error) {
	// Obtiene los roles desde el repositorio
	result, err := p.repository.Get(ctx, map[string]interface{}{}, nil, nil)
	if err != nil {
		return nil, err
	}

	// Convierte los resultados a la estructura de respuesta
	resp = make([]models.ProgramResp, len(result))
	for i, v := range result {
		// Asegúrate de que el tipo sea correcto antes de hacer la conversión
		if program, ok := v.(models.ProgramResp); ok {
			resp[i] = models.ProgramResp(program) // Asumiendo que tienes un constructor o un mapeo en tu struct
		} else {
			return nil, wrappers.NewNonExistentErr(fmt.Errorf("error de tipo: se esperaba entities.Products, pero se obtuvo %T", v))
		}
	}
	return
}

// GetByID user
func (p *programService) GetByID(ctx context.Context, ID string) (resp models.ProgramResp, err error) {
	program, err := p.repository.GetByID(ctx, ID)

	if err != nil {
		return resp, fmt.Errorf("programa con ID %s no encontrado", ID)
	}

	if program == nil {
		// Si no se encuentra el colegio (colegios es nil), devolver un valor en blanco y un error
		return models.ProgramResp{}, fmt.Errorf("programa con ID %s no encontrado", ID)
	}

	resp = *program.(*models.ProgramResp)

	return
}

// Update user
func (p *programService) Update(ctx context.Context, ID string, program models.UpdateProgramReq) (err error) {

	dbProgram, err := p.GetByID(ctx, ID)
	if err != nil {
		return
	}
	// Actualizar los campos solo si no son nil
	// Actualizar la fecha de modificación
	if program.Name != nil {
		dbProgram.Name = *program.Name
	}
	if program.Valor1 != nil {
		dbProgram.Valor1 = *program.Valor1
	}
	if program.Valor2 != nil {
		dbProgram.Valor2 = *program.Valor2
	}
	if program.Valor3 != nil {
		dbProgram.Valor3 = *program.Valor3
	}
	if program.Valor4 != nil {
		dbProgram.Valor4 = *program.Valor4
	}
	if program.Valor5 != nil {
		dbProgram.Valor5 = *program.Valor5
	}
	if program.Active != nil {
		dbProgram.Active = *program.Active
	}
	if program.Reserva != nil {
		dbProgram.Reserva = *program.Reserva
	}
	if program.Author != nil {
		dbProgram.Author = *program.Author
	}

	// Llamar al repositorio para actualizar la entidad
	err = p.repository.Update(ctx, ID, models.Program(dbProgram))

	return err
}

// Delete user
func (p *programService) Delete(ctx context.Context, ID string) (err error) {
	err = p.repository.Delete(ctx, ID)
	if errors.Is(err, wrappers.NonExistentErr) {
		err = wrappers.NewNonExistentErr(fmt.Errorf("ID %s not found", ID))
	}

	return err
}
