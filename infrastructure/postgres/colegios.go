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
type colegiosRepository struct {
	infrastructure.PostgresRepositoryOrm
}

// NewUserRepository creates a roles repository for postgres
func NewColegiosRepository(ctx context.Context, db *gorm.DB) ports.ColegiosRepository {
	return &colegiosRepository{
		infrastructure.PostgresRepositoryOrm{
			DB: db,
		},
	}
}

func (s *colegiosRepository) Create(ctx context.Context, colegios interface{}) (string, error) {

	// Asegúrate de que el tipo del usuario es correcto
	u := colegios.(models.CreateColegiosReq)

	var existingColegios = colegios.(models.CreateColegiosReq)

	err := s.DB.Table("establecimiento").Model(&existingColegios).Where("codigo = ?", u.Codigo).First(&existingColegios).Error
	if err == nil {
		// Si no hay error, significa que se encontró un rol con ese nombre
		return "error", errors.New("El Colegio con el Codigo '" + u.Codigo + "' ya existe")
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		// Si el error no es de registro no encontrado, es un error inesperado
		return "error", errors.New("Error al buscar colegio: " + err.Error())
	}

	// Usamos el contexto y creamos el registro en la base de datos
	if err := s.DB.WithContext(ctx).Table("establecimiento").Create(&u).Error; err != nil {
		return "", err
	}

	// El ID se asigna automáticamente a la estructura 'u' después de la creación
	return u.ID, nil
}

func (s *colegiosRepository) Get(ctx context.Context, filter map[string]interface{}, skip, take *int) ([]interface{}, error) {
	var colegios []models.ColegiosResp

	// Crea una consulta base
	query := s.DB.WithContext(ctx).Model(&models.ColegiosResp{})

	// Aplica filtros si existen
	for key, value := range filter {
		query = query.Where(fmt.Sprintf("%s = ?", key), value)
	}

	// Aplica paginación si 'skip' y 'take' están definidos
	if skip != nil && take != nil {
		query = query.Offset(*skip).Limit(*take)
	}

	// Ejecuta la consulta
	if err := query.Table("establecimiento").Order("id ASC").Find(&colegios).Error; err != nil {
		return nil, err
	}

	if len(colegios) < 1 {
		return nil, wrappers.NewNonExistentErr(sql.ErrNoRows)
	}

	// Convierte []entities.Products a []interface{}
	result := make([]interface{}, len(colegios))
	for i, company := range colegios {
		result[i] = company
	}

	return result, nil
}

func (s *colegiosRepository) GetByID(ctx context.Context, ID string) (interface{}, error) {

	var colegios models.ColegiosResp

	// Busca el registro en la base de datos utilizando GORM
	if err := s.DB.WithContext(ctx).Table("establecimiento").Where("id = ?", ID).First(&colegios).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, wrappers.NewNonExistentErr(err)
		}
		return nil, err
	}

	return &colegios, nil
}

func (s *colegiosRepository) Update(ctx context.Context, ID string, colegios interface{}) error {

	if err := s.DB.WithContext(ctx).Table("establecimiento").Model(&models.UpdateColegiosReq{}).Where("id = ?", ID).Updates(colegios).Error; err != nil {
		if strings.Contains(err.Error(), "record not found") {
			return wrappers.NewNonExistentErr(err)
		} else {
			return err
		}
	}

	return nil
}

func (s *colegiosRepository) Delete(ctx context.Context, ID string) error {
	var colegios models.CreateColegiosReq

	result := s.DB.WithContext(ctx).Table("establecimiento").Where("id = ?", ID).Delete(&colegios)

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
