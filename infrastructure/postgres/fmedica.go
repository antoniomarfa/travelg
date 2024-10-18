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
type fmedicaRepository struct {
	infrastructure.PostgresRepositoryOrm
}

// NewUserRepository creates a roles repository for postgres
func NewFmedicaRepository(ctx context.Context, db *gorm.DB) ports.FmedicaRepository {
	return &fmedicaRepository{
		infrastructure.PostgresRepositoryOrm{
			DB: db,
		},
	}
}

func (s *fmedicaRepository) Create(ctx context.Context, ficha interface{}) (string, error) {

	// Asegúrate de que el tipo del usuario es correcto
	u := ficha.(models.CreateFmedicaReq)

	var existingFicha = ficha.(models.CreateFmedicaReq)

	err := s.DB.Model(&existingFicha).Where("rutalumn = ?", u.Rutalumn).First(&existingFicha).Error
	if err == nil {
		// Si no hay error, significa que se encontró un rol con ese nombre
		return "error", errors.New("La ficha medica del alumno '" + u.Rutalumn + "' ya existe")
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		// Si el error no es de registro no encontrado, es un error inesperado
		return "error", errors.New("Error al buscar ficha medica: " + err.Error())
	}

	// Usamos el contexto y creamos el registro en la base de datos
	if err := s.DB.WithContext(ctx).Create(&u).Error; err != nil {
		return "", err
	}

	// El ID se asigna automáticamente a la estructura 'u' después de la creación
	return u.ID, nil
}

func (s *fmedicaRepository) Get(ctx context.Context, filter map[string]interface{}, skip, take *int) ([]interface{}, error) {
	var ficha []models.FmedicaResp

	// Crea una consulta base
	query := s.DB.WithContext(ctx).Model(&models.FmedicaResp{})

	// Aplica filtros si existen
	for key, value := range filter {
		query = query.Where(fmt.Sprintf("%s = ?", key), value)
	}

	// Aplica paginación si 'skip' y 'take' están definidos
	if skip != nil && take != nil {
		query = query.Offset(*skip).Limit(*take)
	}

	// Ejecuta la consulta
	if err := query.Table("fichamedica").Find(&ficha).Error; err != nil {
		return nil, err
	}

	if len(ficha) < 1 {
		return nil, wrappers.NewNonExistentErr(sql.ErrNoRows)
	}

	// Convierte []entities.Products a []interface{}
	result := make([]interface{}, len(ficha))
	for i, ficha := range ficha {
		result[i] = ficha
	}

	return result, nil
}

func (s *fmedicaRepository) GetByID(ctx context.Context, ID string) (interface{}, error) {

	var ficha models.FmedicaResp

	// Busca el registro en la base de datos utilizando GORM
	if err := s.DB.WithContext(ctx).Table("fichamedica").Where("id = ?", ID).First(&ficha).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, wrappers.NewNonExistentErr(err)
		}
		return nil, wrappers.NewNonExistentErr(err)
	}

	return &ficha, nil
}

func (s *fmedicaRepository) Update(ctx context.Context, ID string, ficha interface{}) error {

	if err := s.DB.WithContext(ctx).Table("fichamedica").Model(&models.UpdateFmedicaReq{}).Where("id = ?", ID).Updates(ficha).Error; err != nil {
		if strings.Contains(err.Error(), "record not found") {
			return wrappers.NewNonExistentErr(err)
		} else {
			return err
		}
	}

	return nil
}

func (s *fmedicaRepository) Delete(ctx context.Context, ID string) error {
	var ficha models.CreateFmedicaReq

	result := s.DB.WithContext(ctx).Table("fichamedica").Where("id = ?", ID).Delete(&ficha)

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
