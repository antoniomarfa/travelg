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
type usersService struct {
	config     config.Config
	repository ports.UsersRepository
}

// NewURolesService creates a new user service
func NewUsersService(cfg config.Config, repo ports.UsersRepository) ports.UsersService {
	return &usersService{
		config:     cfg,
		repository: repo,
	}
}

// Create roles
func (p *usersService) Create(ctx context.Context, users models.CreateUsersReq) (resp models.CreationResp, err error) {

	insertedID, err := p.repository.Create(ctx, models.CreateUsersReq(users))
	if err != nil {
		return
	}

	resp = models.CreationResp{
		InsertedID: insertedID,
	}

	return
}

// GetAll users
func (p *usersService) GetAll(ctx context.Context) (resp []models.UsersResp, err error) {
	// Obtiene los roles desde el repositorio
	result, err := p.repository.Get(ctx, map[string]interface{}{}, nil, nil)
	if err != nil {
		return nil, err
	}

	// Convierte los resultados a la estructura de respuesta
	resp = make([]models.UsersResp, len(result))
	for i, v := range result {
		// Asegúrate de que el tipo sea correcto antes de hacer la conversión
		if users, ok := v.(models.UsersResp); ok {
			resp[i] = models.UsersResp(users) // Asumiendo que tienes un constructor o un mapeo en tu struct
		} else {
			return nil, wrappers.NewNonExistentErr(fmt.Errorf("error de tipo: se esperaba entities.Products, pero se obtuvo %T", v))
		}
	}
	return
}

// GetByID user
func (p *usersService) GetByID(ctx context.Context, ID string) (resp models.UsersResp, err error) {
	users, err := p.repository.GetByID(ctx, ID)

	if err != nil {
		return resp, fmt.Errorf("usuario con ID %s no encontrado", ID)
	}

	if users == nil {
		// Si no se encuentra el colegio (colegios es nil), devolver un valor en blanco y un error
		return models.UsersResp{}, fmt.Errorf("usuario con ID %s no encontrado", ID)
	}

	resp = *users.(*models.UsersResp)

	return
}

// Update user
func (p *usersService) Update(ctx context.Context, ID string, users models.UpdateUsersReq) (err error) {

	dbUsers, err := p.GetByID(ctx, ID)
	if err != nil {
		return
	}
	// Actualizar los campos solo si no son nil
	// Actualizar la fecha de modificación
	now := time.Now().UTC()
	if users.Username != nil {
		dbUsers.Username = *users.Username
	}

	if users.Name != nil {
		dbUsers.Name = *users.Name
	}
	if users.Email != nil {
		dbUsers.Email = *users.Email
	}
	if users.Phone != nil {
		dbUsers.Phone = *users.Phone
	}
	if users.RolesId != nil {
		dbUsers.RolesId = *users.RolesId
	}
	if users.Active != nil {
		dbUsers.Active = *users.Active
	}
	if users.Author != nil {
		dbUsers.Author = *users.Author
	}

	dbUsers.UpdatedDate = now
	// Llamar al repositorio para actualizar la entidad
	err = p.repository.Update(ctx, ID, models.Users(dbUsers))

	return err
}

// Delete user
func (p *usersService) Delete(ctx context.Context, ID string) (err error) {
	err = p.repository.Delete(ctx, ID)
	if errors.Is(err, wrappers.NonExistentErr) {
		err = wrappers.NewNonExistentErr(fmt.Errorf("ID %s not found", ID))
	}

	return err
}
