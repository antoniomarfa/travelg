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
type pagosRepository struct {
	infrastructure.PostgresRepositoryOrm
}

// NewUserRepository creates a roles repository for postgres
func NewPagosRepository(ctx context.Context, db *gorm.DB) ports.PagosRepository {
	return &pagosRepository{
		infrastructure.PostgresRepositoryOrm{
			DB: db,
		},
	}
}

func (s *pagosRepository) Create(ctx context.Context, pagos interface{}) (string, error) {

	// Asegúrate de que el tipo del usuario es correcto
	u := pagos.(models.CreatePagosReq)

	var existingPagos = pagos.(models.CreatePagosReq)

	err := s.DB.Model(&existingPagos).Where("identificador = ?", u.Identificador).First(&existingPagos).Error
	if err == nil {
		// Si no hay error, significa que se encontró un rol con ese nombre
		return "error", errors.New("El Ingreso con el identificador '" + u.Identificador + "' ya existe")
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		// Si el error no es de registro no encontrado, es un error inesperado
		return "error", errors.New("Error al buscar ingreso: " + err.Error())
	}

	// Usamos el contexto y creamos el registro en la base de datos
	if err := s.DB.WithContext(ctx).Create(&u).Error; err != nil {
		return "", err
	}

	// El ID se asigna automáticamente a la estructura 'u' después de la creación
	return u.ID, nil
}

func (s *pagosRepository) Get(ctx context.Context, filter map[string]interface{}, skip, take *int) ([]interface{}, error) {
	var pagos []models.PagosResp

	// Crea una consulta base
	query := s.DB.WithContext(ctx).Model(&models.PagosResp{})

	// Aplica filtros si existen
	for key, value := range filter {
		query = query.Where(fmt.Sprintf("%s = ?", key), value)
	}

	// Aplica paginación si 'skip' y 'take' están definidos
	if skip != nil && take != nil {
		query = query.Offset(*skip).Limit(*take)
	}

	// Ejecuta la consulta
	if err := query.Table("pagos").Find(&pagos).Error; err != nil {
		return nil, err
	}

	if len(pagos) < 1 {
		return nil, wrappers.NewNonExistentErr(sql.ErrNoRows)
	}

	// Convierte []entities.Products a []interface{}
	result := make([]interface{}, len(pagos))
	for i, pagos := range pagos {
		result[i] = pagos
	}

	return result, nil
}

func (s *pagosRepository) GetByID(ctx context.Context, ID string) (interface{}, error) {

	var pagos models.PagosResp

	// Busca el registro en la base de datos utilizando GORM
	if err := s.DB.WithContext(ctx).Table("pagos").Where("id = ?", ID).First(&pagos).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, wrappers.NewNonExistentErr(err)
		}
		return nil, wrappers.NewNonExistentErr(err)
	}

	return &pagos, nil
}

func (s *pagosRepository) Update(ctx context.Context, ID string, pagos interface{}) error {

	if err := s.DB.WithContext(ctx).Table("pagos").Model(&models.UpdatePagosReq{}).Where("id = ?", ID).Updates(pagos).Error; err != nil {
		if strings.Contains(err.Error(), "record not found") {
			return wrappers.NewNonExistentErr(err)
		} else {
			return err
		}
	}

	return nil
}

func (s *pagosRepository) Delete(ctx context.Context, ID string) error {
	var pagos models.CreatePagosReq

	result := s.DB.WithContext(ctx).Table("pagos").Where("id = ?", ID).Delete(&pagos)

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
