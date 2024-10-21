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
type rolesRepository struct {
	infrastructure.PostgresRepositoryOrm
}

// NewUserRepository creates a roles repository for postgres
func NewRolesRepository(ctx context.Context, db *gorm.DB) ports.RolesRepository {
	return &rolesRepository{
		infrastructure.PostgresRepositoryOrm{
			DB: db,
		},
	}
}

func (s *rolesRepository) Create(ctx context.Context, roles interface{}) (string, error) {

	// Asegúrate de que el tipo del usuario es correcto
	u := roles.(models.CreateRolesReq)

	var existingReg = roles.(models.CreateRolesReq)

	err := s.DB.Model(&existingReg).Where("description = ?", u.Description).First(&existingReg).Error
	if err == nil {
		// Si no hay error, significa que se encontró un rol con ese nombre
		return "error", errors.New("El Role con la descripcion '" + u.Description + "' ya existe")
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		// Si el error no es de registro no encontrado, es un error inesperado
		return "error", errors.New("Error al buscar roles: " + err.Error())
	}

	// Usamos el contexto y creamos el registro en la base de datos
	if err := s.DB.WithContext(ctx).Create(&u).Error; err != nil {
		return "", err
	}

	// El ID se asigna automáticamente a la estructura 'u' después de la creación
	return u.ID, nil
}

func (s *rolesRepository) Get(ctx context.Context, filter map[string]interface{}, skip, take *int) ([]interface{}, error) {
	var registro []models.RolesResp

	// Crea una consulta base
	query := s.DB.WithContext(ctx).Model(&models.RolesResp{})

	// Aplica filtros si existen
	for key, value := range filter {
		query = query.Where(fmt.Sprintf("%s = ?", key), value)
	}

	// Aplica paginación si 'skip' y 'take' están definidos
	if skip != nil && take != nil {
		query = query.Offset(*skip).Limit(*take)
	}

	// Ejecuta la consulta
	if err := query.Table("roles").Order("id asc").Find(&registro).Error; err != nil {
		return nil, err
	}

	if len(registro) < 1 {
		return nil, wrappers.NewNonExistentErr(sql.ErrNoRows)
	}

	// Convierte []entities.Products a []interface{}
	result := make([]interface{}, len(registro))
	for i, roles := range registro {
		result[i] = roles
	}

	return result, nil
}

func (s *rolesRepository) GetByID(ctx context.Context, ID string) (interface{}, error) {

	var registro models.RolesResp

	// Busca el registro en la base de datos utilizando GORM
	if err := s.DB.WithContext(ctx).Table("roles").Where("id = ?", ID).First(&registro).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, wrappers.NewNonExistentErr(err)
		}
		return nil, wrappers.NewNonExistentErr(err)
	}

	return &registro, nil
}

func (s *rolesRepository) Update(ctx context.Context, ID string, roles interface{}) error {

	if err := s.DB.WithContext(ctx).Table("roles").Model(&models.UpdateRolesReq{}).Where("id = ?", ID).Updates(roles).Error; err != nil {
		if strings.Contains(err.Error(), "record not found") {
			return wrappers.NewNonExistentErr(err)
		} else {
			return err
		}
	}

	return nil
}

func (s *rolesRepository) Delete(ctx context.Context, ID string) error {
	var roles models.CreateRolesReq

	result := s.DB.WithContext(ctx).Table("roles").Where("id = ?", ID).Delete(&roles)

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
