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
type comunasRepository struct {
	infrastructure.PostgresRepositoryOrm
}

// NewUserRepository creates a roles repository for postgres
func NewComunasRepository(ctx context.Context, db *gorm.DB) ports.ComunasRepository {
	return &comunasRepository{
		infrastructure.PostgresRepositoryOrm{
			DB: db,
		},
	}
}

func (s *comunasRepository) Create(ctx context.Context, comunas interface{}) (string, error) {

	// Asegúrate de que el tipo del usuario es correcto
	u := comunas.(models.CreateComunasReq)

	var existingComunas = comunas.(models.CreateComunasReq)

	err := s.DB.Model(&existingComunas).Where("description = ?", u.Description).First(&existingComunas).Error
	if err == nil {
		// Si no hay error, significa que se encontró un rol con ese nombre
		return "error", errors.New("la comuna con la descripcion '" + u.Description + "' ya existe")
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		// Si el error no es de registro no encontrado, es un error inesperado
		return "error", errors.New("Error al buscar comuna: " + err.Error())
	}

	// Usamos el contexto y creamos el registro en la base de datos
	if err := s.DB.WithContext(ctx).Create(&u).Error; err != nil {
		return "", err
	}

	// El ID se asigna automáticamente a la estructura 'u' después de la creación
	return u.ID, nil
}

func (s *comunasRepository) Get(ctx context.Context, filter map[string]interface{}, skip, take *int) ([]interface{}, error) {
	var comunas []models.ComunasResp

	// Crea una consulta base
	query := s.DB.WithContext(ctx).Model(&models.ComunasResp{})

	// Aplica filtros si existen
	for key, value := range filter {
		query = query.Where(fmt.Sprintf("%s = ?", key), value)
	}

	// Aplica paginación si 'skip' y 'take' están definidos
	if skip != nil && take != nil {
		query = query.Offset(*skip).Limit(*take)
	}

	// Ejecuta la consulta
	if err := query.Table("communes").Find(&comunas).Error; err != nil {
		return nil, err
	}

	if len(comunas) < 1 {
		return nil, wrappers.NewNonExistentErr(sql.ErrNoRows)
	}

	// Convierte []entities.Products a []interface{}
	result := make([]interface{}, len(comunas))
	for i, comunas := range comunas {
		result[i] = comunas
	}

	return result, nil
}

func (s *comunasRepository) GetByID(ctx context.Context, ID string) (interface{}, error) {

	var comunas models.ComunasResp

	// Busca el registro en la base de datos utilizando GORM
	if err := s.DB.WithContext(ctx).Table("communes").Where("id = ?", ID).First(&comunas).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, wrappers.NewNonExistentErr(err)
		}
		return nil, wrappers.NewNonExistentErr(err)
	}

	return &comunas, nil
}

func (s *comunasRepository) Update(ctx context.Context, ID string, comunas interface{}) error {

	if err := s.DB.WithContext(ctx).Table("communes").Model(&models.UpdateComunasReq{}).Where("id = ?", ID).Updates(comunas).Error; err != nil {
		if strings.Contains(err.Error(), "record not found") {
			return wrappers.NewNonExistentErr(err)
		} else {
			return err
		}
	}

	return nil
}

func (s *comunasRepository) Delete(ctx context.Context, ID string) error {
	var comunas models.CreateComunasReq

	result := s.DB.WithContext(ctx).Table("communes").Where("id = ?", ID).Delete(&comunas)

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
