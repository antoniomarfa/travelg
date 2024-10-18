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
type companyRepository struct {
	infrastructure.PostgresRepositoryOrm
}

// NewUserRepository creates a roles repository for postgres
func NewCompanyRepository(ctx context.Context, db *gorm.DB) ports.CompanyRepository {
	return &companyRepository{
		infrastructure.PostgresRepositoryOrm{
			DB: db,
		},
	}
}

func (s *companyRepository) Create(ctx context.Context, company interface{}) (string, error) {

	// Asegúrate de que el tipo del usuario es correcto
	u := company.(models.CreateCompanyReq)

	var existingCompany = company.(models.CreateCompanyReq)

	err := s.DB.Model(&existingCompany).Where("rut = ?", u.Rut).First(&existingCompany).Error
	if err == nil {
		// Si no hay error, significa que se encontró un rol con ese nombre
		return "error", errors.New("El Cliente con el Rut '" + u.Rut + "' ya existe")
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		// Si el error no es de registro no encontrado, es un error inesperado
		return "error", errors.New("Error al buscar cliente: " + err.Error())
	}

	// Usamos el contexto y creamos el registro en la base de datos
	if err := s.DB.WithContext(ctx).Create(&u).Error; err != nil {
		return "", err
	}

	// El ID se asigna automáticamente a la estructura 'u' después de la creación
	return u.ID, nil
}

func (s *companyRepository) Get(ctx context.Context, filter map[string]interface{}, skip, take *int) ([]interface{}, error) {
	var company []models.CompanyResp

	// Crea una consulta base
	query := s.DB.WithContext(ctx).Model(&models.CompanyResp{})

	// Aplica filtros si existen
	for key, value := range filter {
		query = query.Where(fmt.Sprintf("%s = ?", key), value)
	}

	// Aplica paginación si 'skip' y 'take' están definidos
	if skip != nil && take != nil {
		query = query.Offset(*skip).Limit(*take)
	}

	// Ejecuta la consulta
	if err := query.Table("company").Find(&company).Error; err != nil {
		return nil, err
	}

	if len(company) < 1 {
		return nil, wrappers.NewNonExistentErr(sql.ErrNoRows)
	}

	// Convierte []entities.Products a []interface{}
	result := make([]interface{}, len(company))
	for i, company := range company {
		result[i] = company
	}

	return result, nil
}

func (s *companyRepository) GetByID(ctx context.Context, ID string) (interface{}, error) {

	var company models.CompanyResp

	// Busca el registro en la base de datos utilizando GORM
	if err := s.DB.WithContext(ctx).Table("company").Where("id = ?", ID).First(&company).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, wrappers.NewNonExistentErr(err)
		}
		return nil, wrappers.NewNonExistentErr(err)
	}

	return &company, nil
}

func (s *companyRepository) Update(ctx context.Context, ID string, company interface{}) error {

	if err := s.DB.WithContext(ctx).Table("company").Model(&models.UpdateCompanyReq{}).Where("id = ?", ID).Updates(company).Error; err != nil {
		if strings.Contains(err.Error(), "record not found") {
			return wrappers.NewNonExistentErr(err)
		} else {
			return err
		}
	}

	return nil
}

func (s *companyRepository) Delete(ctx context.Context, ID string) error {
	var company models.CreateCompanyReq

	result := s.DB.WithContext(ctx).Table("company").Where("id = ?", ID).Delete(&company)

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
