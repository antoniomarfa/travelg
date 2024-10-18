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
type cursoRepository struct {
	infrastructure.PostgresRepositoryOrm
}

// NewUserRepository creates a roles repository for postgres
func NewCursoRepository(ctx context.Context, db *gorm.DB) ports.CursoRepository {
	return &cursoRepository{
		infrastructure.PostgresRepositoryOrm{
			DB: db,
		},
	}
}

func (s *cursoRepository) Create(ctx context.Context, curso interface{}) (string, error) {

	// Asegúrate de que el tipo del usuario es correcto
	u := curso.(models.CreateCursoReq)

	var existingCurso = curso.(models.CreateCursoReq)

	err := s.DB.Model(&existingCurso).Where("rutalumno = ?", u.Rutalumno).First(&existingCurso).Error
	if err == nil {
		// Si no hay error, significa que se encontró un rol con ese nombre
		return "error", errors.New("El alumno '" + u.Rutalumno + "' ya existe")
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		// Si el error no es de registro no encontrado, es un error inesperado
		return "error", errors.New("Error al buscar alumnos: " + err.Error())
	}

	// Usamos el contexto y creamos el registro en la base de datos
	if err := s.DB.WithContext(ctx).Create(&u).Error; err != nil {
		return "", err
	}

	// El ID se asigna automáticamente a la estructura 'u' después de la creación
	return u.ID, nil
}

func (s *cursoRepository) Get(ctx context.Context, filter map[string]interface{}, skip, take *int) ([]interface{}, error) {
	var curso []models.CursoResp

	// Crea una consulta base
	query := s.DB.WithContext(ctx).Model(&models.CursoResp{})

	// Aplica filtros si existen
	for key, value := range filter {
		query = query.Where(fmt.Sprintf("%s = ?", key), value)
	}
	/*
		// Aplica paginación si 'skip' y 'take' están definidos
		if skip != nil && take != nil {
			query = query.Offset(*skip).Limit(*take)
		}
	*/
	// Ejecuta la consulta
	if err := query.Table("curso").Find(&curso).Error; err != nil {
		return nil, err
	}

	if len(curso) < 1 {
		return nil, wrappers.NewNonExistentErr(sql.ErrNoRows)
	}

	// Convierte []entities.Products a []interface{}
	result := make([]interface{}, len(curso))
	for i, curso := range curso {
		result[i] = curso
	}

	return result, nil
}

func (s *cursoRepository) GetByID(ctx context.Context, ID string) (interface{}, error) {

	var curso models.CursoResp

	// Busca el registro en la base de datos utilizando GORM
	if err := s.DB.WithContext(ctx).Table("curso").Where("id = ?", ID).First(&curso).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, wrappers.NewNonExistentErr(err)
		}
		return nil, wrappers.NewNonExistentErr(err)
	}

	return &curso, nil
}

func (s *cursoRepository) Update(ctx context.Context, ID string, curso interface{}) error {

	if err := s.DB.WithContext(ctx).Table("curso").Model(&models.UpdateCursoReq{}).Where("id = ?", ID).Updates(curso).Error; err != nil {
		if strings.Contains(err.Error(), "record not found") {
			return wrappers.NewNonExistentErr(err)
		} else {
			return err
		}
	}

	return nil
}

func (s *cursoRepository) Delete(ctx context.Context, ID string) error {
	var curso models.CreateCursoReq

	result := s.DB.WithContext(ctx).Table("curso").Where("id = ?", ID).Delete(&curso)

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
