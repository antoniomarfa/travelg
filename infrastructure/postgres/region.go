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
type regionRepository struct {
	infrastructure.PostgresRepositoryOrm
}

// NewUserRepository creates a roles repository for postgres
func NewRegionRepository(ctx context.Context, db *gorm.DB) ports.RegionRepository {
	return &regionRepository{
		infrastructure.PostgresRepositoryOrm{
			DB: db,
		},
	}
}

func (s *regionRepository) Create(ctx context.Context, region interface{}) (string, error) {

	// Asegúrate de que el tipo del usuario es correcto
	u := region.(models.CreateRegionReq)

	var existingReg = region.(models.CreateRegionReq)

	err := s.DB.Model(&existingReg).Where("code = ?", u.Code).First(&existingReg).Error
	if err == nil {
		// Si no hay error, significa que se encontró un rol con ese nombre
		return "error", errors.New("La Region con el codigo '" + u.Code + "' ya existe")
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		// Si el error no es de registro no encontrado, es un error inesperado
		return "error", errors.New("Error al buscar region: " + err.Error())
	}

	// Usamos el contexto y creamos el registro en la base de datos
	if err := s.DB.WithContext(ctx).Create(&u).Error; err != nil {
		return "", err
	}

	// El ID se asigna automáticamente a la estructura 'u' después de la creación
	return u.ID, nil
}

func (s *regionRepository) Get(ctx context.Context, filter map[string]interface{}, skip, take *int) ([]interface{}, error) {
	var registro []models.RegionResp

	// Crea una consulta base
	query := s.DB.WithContext(ctx).Model(&models.RegionResp{})

	// Aplica filtros si existen
	for key, value := range filter {
		query = query.Where(fmt.Sprintf("%s = ?", key), value)
	}

	// Aplica paginación si 'skip' y 'take' están definidos
	if skip != nil && take != nil {
		query = query.Offset(*skip).Limit(*take)
	}

	// Ejecuta la consulta
	if err := query.Table("region").Order("id asc").Find(&registro).Error; err != nil {
		return nil, err
	}

	if len(registro) < 1 {
		return nil, wrappers.NewNonExistentErr(sql.ErrNoRows)
	}

	// Convierte []entities.Products a []interface{}
	result := make([]interface{}, len(registro))
	for i, region := range registro {
		result[i] = region
	}

	return result, nil
}

func (s *regionRepository) GetByID(ctx context.Context, ID string) (interface{}, error) {

	var registro models.RegionResp

	// Busca el registro en la base de datos utilizando GORM
	if err := s.DB.WithContext(ctx).Table("region").Where("id = ?", ID).First(&registro).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, wrappers.NewNonExistentErr(err)
		}
		return nil, wrappers.NewNonExistentErr(err)
	}

	return &registro, nil
}

func (s *regionRepository) Update(ctx context.Context, ID string, region interface{}) error {

	if err := s.DB.WithContext(ctx).Table("region").Model(&models.UpdateRegionReq{}).Where("id = ?", ID).Updates(region).Error; err != nil {
		if strings.Contains(err.Error(), "record not found") {
			return wrappers.NewNonExistentErr(err)
		} else {
			return err
		}
	}

	return nil
}

func (s *regionRepository) Delete(ctx context.Context, ID string) error {
	var region models.CreateRegionReq

	result := s.DB.WithContext(ctx).Table("region").Where("id = ?", ID).Delete(&region)

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
