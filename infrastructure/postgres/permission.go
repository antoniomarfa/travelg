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
type permissionRepository struct {
	infrastructure.PostgresRepositoryOrm
}

// NewUserRepository creates a roles repository for postgres
func NewPermissionRepository(ctx context.Context, db *gorm.DB) ports.PermissionRepository {
	return &permissionRepository{
		infrastructure.PostgresRepositoryOrm{
			DB: db,
		},
	}
}

func (s *permissionRepository) Create(ctx context.Context, permission interface{}) (string, error) {

	// Asegúrate de que el tipo del usuario es correcto
	u := permission.(models.CreateRolesPermissionsReq)
	/*
		var existingReg = permission.(models.CreateRolesPermissionsReq)

		err := s.DB.Table("roles_permissions").Model(&existingReg).Where("id = ?", u.ID).First(&existingReg).Error
		if err == nil {
			// Si no hay error, significa que se encontró un rol con ese nombre
			return "error", errors.New("El Permission con el codigo '" + u.ID + "' ya existe")
		}
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			// Si el error no es de registro no encontrado, es un error inesperado
			return "error", errors.New("Error al buscar permission: " + err.Error())
		}
	*/
	// Usamos el contexto y creamos el registro en la base de datos
	if err := s.DB.WithContext(ctx).Table("roles_permissions").Create(&u).Error; err != nil {
		return "", err
	}

	// El ID se asigna automáticamente a la estructura 'u' después de la creación
	return u.ID, nil
}

func (s *permissionRepository) Get(ctx context.Context, filter map[string]interface{}, skip, take *int) ([]interface{}, error) {
	var registro []models.RolesPermissionsResp

	// Crea una consulta base
	query := s.DB.WithContext(ctx).Model(&models.ProgramResp{})

	// Aplica filtros si existen
	for key, value := range filter {
		query = query.Where(fmt.Sprintf("%s = ?", key), value)
	}

	// Aplica paginación si 'skip' y 'take' están definidos
	if skip != nil && take != nil {
		query = query.Offset(*skip).Limit(*take)
	}

	// Ejecuta la consulta
	if err := query.Preload("Role").Table("roles_permissions").Find(&registro).Error; err != nil {
		return nil, err
	}

	if len(registro) < 1 {
		return nil, wrappers.NewNonExistentErr(sql.ErrNoRows)
	}

	// Convierte []entities.Products a []interface{}
	result := make([]interface{}, len(registro))
	for i, permission := range registro {
		result[i] = permission
	}

	return result, nil
}

func (s *permissionRepository) GetByID(ctx context.Context, ID string) (interface{}, error) {

	var registro models.RolesPermissionsResp

	// Busca el registro en la base de datos utilizando GORM
	if err := s.DB.WithContext(ctx).Table("roles_permission").Where("id = ?", ID).First(&registro).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, wrappers.NewNonExistentErr(err)
		}
		return nil, wrappers.NewNonExistentErr(err)
	}

	return &registro, nil
}

func (s *permissionRepository) Update(ctx context.Context, ID string, permission interface{}) error {

	if err := s.DB.WithContext(ctx).Table("roles_permissions").Model(&models.UpdateRolesPermissionsReq{}).Where("id = ?", ID).Updates(permission).Error; err != nil {
		if strings.Contains(err.Error(), "record not found") {
			return wrappers.NewNonExistentErr(err)
		} else {
			return err
		}
	}

	return nil
}

func (s *permissionRepository) Delete(ctx context.Context, ID string) error {
	var permission models.CreateRolesPermissionsReq

	result := s.DB.WithContext(ctx).Table("roles_permissions").Where("roles_id = ?", ID).Delete(&permission)

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
