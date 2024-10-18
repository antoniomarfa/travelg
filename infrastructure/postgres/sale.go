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
type saleRepository struct {
	infrastructure.PostgresRepositoryOrm
}

// NewUserRepository creates a roles repository for postgres
func NewSaleRepository(ctx context.Context, db *gorm.DB) ports.SaleRepository {
	return &saleRepository{
		infrastructure.PostgresRepositoryOrm{
			DB: db,
		},
	}
}

func (s *saleRepository) Create(ctx context.Context, sale interface{}) (string, error) {

	// Asegúrate de que el tipo del usuario es correcto
	u := sale.(models.CreateSaleReq)

	var existingSale = sale.(models.CreateSaleReq)

	err := s.DB.Table("sale").Model(&existingSale).Where("establecimiento_id = ?", u.EstablecimientoId).First(&existingSale).Error
	if err == nil {
		// Si no hay error, significa que se encontró un rol con ese nombre
		return "error", errors.New("Hay una venta para este colegio, curso ")
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		// Si el error no es de registro no encontrado, es un error inesperado
		return "error", errors.New("Error al buscar ventas: " + err.Error())
	}

	// Usamos el contexto y creamos el registro en la base de datos
	if err := s.DB.WithContext(ctx).Table("sale").Create(&u).Error; err != nil {
		return "", err
	}

	// El ID se asigna automáticamente a la estructura 'u' después de la creación
	return u.ID, nil
}

func (s *saleRepository) Get(ctx context.Context, filter map[string]interface{}, skip, take *int) ([]interface{}, error) {
	var sale []models.SaleResp

	// Crea una consulta base
	query := s.DB.WithContext(ctx).Model(&models.SaleResp{})

	// Aplica filtros si existen
	for key, value := range filter {
		query = query.Where(fmt.Sprintf("%s = ?", key), value)
	}

	// Aplica paginación si 'skip' y 'take' están definidos
	if skip != nil && take != nil {
		query = query.Offset(*skip).Limit(*take)
	}

	// Ejecuta la consulta
	if err := query.Table("sale").Order("id ASC").Find(&sale).Error; err != nil {
		return nil, err
	}

	if len(sale) < 1 {
		return nil, wrappers.NewNonExistentErr(sql.ErrNoRows)
	}

	// Convierte []entities.Products a []interface{}
	result := make([]interface{}, len(sale))
	for i, company := range sale {
		result[i] = company
	}

	return result, nil
}

func (s *saleRepository) GetByID(ctx context.Context, ID string) (interface{}, error) {

	var sale models.SaleResp

	// Busca el registro en la base de datos utilizando GORM
	if err := s.DB.WithContext(ctx).Table("sale").Where("id = ?", ID).First(&sale).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, wrappers.NewNonExistentErr(err)
		}
		return nil, err
	}

	return &sale, nil
}

func (s *saleRepository) Update(ctx context.Context, ID string, sale interface{}) error {

	if err := s.DB.WithContext(ctx).Table("sale").Model(&models.UpdateSaleReq{}).Where("id = ?", ID).Updates(sale).Error; err != nil {
		if strings.Contains(err.Error(), "record not found") {
			return wrappers.NewNonExistentErr(err)
		} else {
			return err
		}
	}

	return nil
}

func (s *saleRepository) Delete(ctx context.Context, ID string) error {
	var sale models.CreateSaleReq

	result := s.DB.WithContext(ctx).Table("sale").Where("id = ?", ID).Delete(&sale)

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
