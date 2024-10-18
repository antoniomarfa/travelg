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
type voucherRepository struct {
	infrastructure.PostgresRepositoryOrm
}

// NewUserRepository creates a roles repository for postgres
func NewVoucherRepository(ctx context.Context, db *gorm.DB) ports.VoucherRepository {
	return &voucherRepository{
		infrastructure.PostgresRepositoryOrm{
			DB: db,
		},
	}
}

func (s *voucherRepository) Create(ctx context.Context, voucher interface{}) (string, error) {

	// Asegúrate de que el tipo del usuario es correcto
	u := voucher.(models.CreateVoucherReq)

	var existingVoucher = voucher.(models.CreateVoucherReq)

	err := s.DB.Model(&existingVoucher).Where("voucher = ?", u.Voucher).First(&existingVoucher).Error
	if err == nil {
		// Si no hay error, significa que se encontró un rol con ese nombre
		return "error", errors.New("El Voucher con el numero '" + u.Voucher + "' ya existe")
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		// Si el error no es de registro no encontrado, es un error inesperado
		return "error", errors.New("Error al buscar voucher: " + err.Error())
	}

	// Usamos el contexto y creamos el registro en la base de datos
	if err := s.DB.WithContext(ctx).Create(&u).Error; err != nil {
		return "", err
	}

	// El ID se asigna automáticamente a la estructura 'u' después de la creación
	return u.ID, nil
}

func (s *voucherRepository) Get(ctx context.Context, filter map[string]interface{}, skip, take *int) ([]interface{}, error) {
	var voucher []models.VoucherResp

	// Crea una consulta base
	query := s.DB.WithContext(ctx).Model(&models.VoucherResp{})

	// Aplica filtros si existen
	for key, value := range filter {
		query = query.Where(fmt.Sprintf("%s = ?", key), value)
	}

	// Aplica paginación si 'skip' y 'take' están definidos
	if skip != nil && take != nil {
		query = query.Offset(*skip).Limit(*take)
	}

	// Ejecuta la consulta
	if err := query.Table("voucher").Find(&voucher).Error; err != nil {
		return nil, err
	}

	if len(voucher) < 1 {
		return nil, wrappers.NewNonExistentErr(sql.ErrNoRows)
	}

	// Convierte []entities.Products a []interface{}
	result := make([]interface{}, len(voucher))
	for i, voucher := range voucher {
		result[i] = voucher
	}

	return result, nil
}

func (s *voucherRepository) GetByID(ctx context.Context, ID string) (interface{}, error) {

	var voucher models.VoucherResp

	// Busca el registro en la base de datos utilizando GORM
	if err := s.DB.WithContext(ctx).Table("voucher").Where("id = ?", ID).First(&voucher).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, wrappers.NewNonExistentErr(err)
		}
		return nil, wrappers.NewNonExistentErr(err)
	}

	return &voucher, nil
}

func (s *voucherRepository) Update(ctx context.Context, ID string, voucher interface{}) error {

	if err := s.DB.WithContext(ctx).Table("voucher").Model(&models.UpdateVoucherReq{}).Where("id = ?", ID).Updates(voucher).Error; err != nil {
		if strings.Contains(err.Error(), "record not found") {
			return wrappers.NewNonExistentErr(err)
		} else {
			return err
		}
	}

	return nil
}

func (s *voucherRepository) Delete(ctx context.Context, ID string) error {
	var voucher models.CreateVoucherReq

	result := s.DB.WithContext(ctx).Table("voucher").Where("id = ?", ID).Delete(&voucher)

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
