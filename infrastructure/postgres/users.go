package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"travel/core/models"
	"travel/core/ports"

	"travel/tools/infrastructure"
	"travel/tools/wrappers"

	"gorm.io/gorm"
)

// userRepository adapter of an roles repository for postgres
type usersRepository struct {
	infrastructure.PostgresRepositoryOrm
}

// NewUserRepository creates a roles repository for postgres
func NewUsersRepository(ctx context.Context, db *gorm.DB) ports.UsersRepository {
	return &usersRepository{
		infrastructure.PostgresRepositoryOrm{
			DB: db,
		},
	}
}

func (s *usersRepository) Create(ctx context.Context, users interface{}) (string, error) {

	// Asegúrate de que el tipo del usuario es correcto
	u := users.(models.CreateUsersReq)

	var existingReg = users.(models.CreateUsersReq)

	err := s.DB.Model(&existingReg).Where("username = ?", u.Username).First(&existingReg).Error
	if err == nil {
		// Si no hay error, significa que se encontró un rol con ese nombre
		return "error", errors.New("El Usuario con el nombre '" + u.Username + "' ya existe")
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		// Si el error no es de registro no encontrado, es un error inesperado
		return "error", errors.New("Error al buscar usuarios: " + err.Error())
	}

	// Usamos el contexto y creamos el registro en la base de datos
	if err := s.DB.WithContext(ctx).Create(&u).Error; err != nil {
		return "", err
	}

	// El ID se asigna automáticamente a la estructura 'u' después de la creación
	return u.ID, nil
}

func (s *usersRepository) Get(ctx context.Context, filter map[string]interface{}, skip, take *int) ([]interface{}, error) {
	var registro []models.UsersResp

	// Crea una consulta base
	query := s.DB.WithContext(ctx).Model(&models.UsersResp{})

	// Aplica filtros si existen
	for key, value := range filter {
		query = query.Where(fmt.Sprintf("%s = ?", key), value)
	}

	// Aplica paginación si 'skip' y 'take' están definidos
	if skip != nil && take != nil {
		query = query.Offset(*skip).Limit(*take)
	}

	// Ejecuta la consulta
	if err := query.Table("users").Find(&registro).Error; err != nil {
		return nil, err
	}

	if len(registro) < 1 {
		return nil, wrappers.NewNonExistentErr(sql.ErrNoRows)
	}

	// Convierte []entities.Products a []interface{}
	result := make([]interface{}, len(registro))
	for i, users := range registro {
		result[i] = users
	}

	return result, nil
}

func (s *usersRepository) GetByID(ctx context.Context, ID string) (interface{}, error) {

	var registro models.UsersResp

	// Busca el registro en la base de datos utilizando GORM
	if err := s.DB.WithContext(ctx).Table("users").Where("id = ?", ID).First(&registro).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, wrappers.NewNonExistentErr(err)
		}
		return nil, wrappers.NewNonExistentErr(err)
	}

	return &registro, nil
}

func (s *usersRepository) Update(ctx context.Context, ID string, users interface{}) error {

	if err := s.DB.WithContext(ctx).Table("roles").Model(&models.UpdateUsersReq{}).Where("id = ?", ID).Updates(users).Error; err != nil {
		if strings.Contains(err.Error(), "record not found") {
			return wrappers.NewNonExistentErr(err)
		} else {
			return err
		}
	}

	return nil
}

func (s *usersRepository) Delete(ctx context.Context, ID string) error {
	var users models.CreateUsersReq

	result := s.DB.WithContext(ctx).Table("users").Where("id = ?", ID).Delete(&users)

	// Comprueba si ocurrió un error durante la eliminación
	if result.Error != nil {
		return wrappers.NewNonExistentErr(result.Error)
	}

	// Comprueba si se eliminaron filas
	if result.RowsAffected == 0 {
		return wrappers.NewNonExistentErr(sql.ErrNoRows) // Manejo de error si no se encontró el registro
	}

	return nil
}
