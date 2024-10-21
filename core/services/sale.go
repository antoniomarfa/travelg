package services

import (
	"context"
	"errors"
	"fmt"

	"travel/config"
	"travel/core/models"
	"travel/core/ports"
	"travel/tools/wrappers"
)

// rolesService adapter of an user service
type saleService struct {
	config     config.Config
	repository ports.CompanyRepository
}

// NewURolesService creates a new user service
func NewSaleService(cfg config.Config, repo ports.SaleRepository) ports.SaleService {
	return &saleService{
		config:     cfg,
		repository: repo,
	}
}

// Create roles
func (p *saleService) Create(ctx context.Context, sale models.CreateSaleReq) (resp models.CreationResp, err error) {

	insertedID, err := p.repository.Create(ctx, models.CreateSaleReq(sale))
	if err != nil {
		return
	}

	resp = models.CreationResp{
		InsertedID: insertedID,
	}

	return
}

// GetAll users
func (p *saleService) GetAll(ctx context.Context, filter map[string]interface{}) (resp []models.SaleResp, err error) {
	// Obtiene los roles desde el repositorio
	result, err := p.repository.Get(ctx, map[string]interface{}{}, nil, nil)
	if err != nil {
		return nil, err
	}

	// Convierte los resultados a la estructura de respuesta
	resp = make([]models.SaleResp, len(result))
	for i, v := range result {
		// Asegúrate de que el tipo sea correcto antes de hacer la conversión
		if sale, ok := v.(models.SaleResp); ok {
			resp[i] = models.SaleResp(sale) // Asumiendo que tienes un constructor o un mapeo en tu struct
		} else {
			return nil, wrappers.NewNonExistentErr(fmt.Errorf("error de tipo: se esperaba, pero se obtuvo %T", v))
		}
	}

	return
}

// GetByID user
func (p *saleService) GetByID(ctx context.Context, ID string) (resp models.SaleResp, err error) {
	sale, err := p.repository.GetByID(ctx, ID)

	if err != nil {
		return resp, fmt.Errorf("venta con ID %s no encontrado", ID)
	}

	if sale == nil {
		// Si no se encuentra el colegio (colegios es nil), devolver un valor en blanco y un error
		return models.SaleResp{}, fmt.Errorf("venta con ID %s no encontrado", ID)
	}

	//	resp = models.ColegiosResp(*colegios.(*models.ColegiosResp))
	resp = *sale.(*models.SaleResp)
	return resp, nil
	// return
}

// Update user
func (p *saleService) Update(ctx context.Context, ID string, sale models.UpdateSaleReq) (err error) {
	dbSale, err := p.GetByID(ctx, ID)
	if err != nil {
		return
	}

	if sale.Fecha != nil {
		dbSale.Fecha = *sale.Fecha
	}
	if sale.SellerId != nil {
		dbSale.SellerId = *sale.SellerId
	}
	if sale.Identificador != nil {
		dbSale.Identificador = *sale.Identificador
	}
	if sale.EstablecimientoId != nil {
		dbSale.EstablecimientoId = *sale.EstablecimientoId
	}
	if sale.ProgramId != nil {
		dbSale.ProgramId = *sale.ProgramId
	}
	if sale.Curso != nil {
		dbSale.Curso = *sale.Curso
	}
	if sale.Idcurso != nil {
		dbSale.Idcurso = *sale.Idcurso
	}
	if sale.Nroalumno != nil {
		dbSale.Nroalumno = *sale.Nroalumno
	}
	if sale.Liberados != nil {
		dbSale.Liberados = *sale.Liberados
	}
	if sale.Program != nil {
		dbSale.Program = *sale.Program
	}
	if sale.Subtotal != nil {
		dbSale.Subtotal = *sale.Subtotal
	}
	if sale.Descm != nil {
		dbSale.Descm = *sale.Descm
	}
	if sale.Vprograma != nil {
		dbSale.Vprograma = *sale.Vprograma
	}
	if sale.Description != nil {
		dbSale.Description = *sale.Description
	}
	if sale.Obs != nil {
		dbSale.Obs = *sale.Obs
	}
	if sale.Fechasalida != nil {
		dbSale.Fechasalida = *sale.Fechasalida
	}
	if sale.Activo != nil {
		dbSale.Activo = *sale.Activo
	}
	if sale.State != nil {
		dbSale.State = *sale.State
	}
	if sale.CorreoEncargado != nil {
		dbSale.CorreoEncargado = *sale.CorreoEncargado
	}
	if sale.Password != nil {
		dbSale.Password = *sale.Password
	}
	if sale.FechaUltpag != nil {
		dbSale.FechaUltpag = *sale.FechaUltpag
	}
	if sale.FechaCierre != nil {
		dbSale.FechaCierre = *sale.FechaCierre
	}
	if sale.Sendemail != nil {
		dbSale.Sendemail = *sale.Sendemail
	}
	if sale.Author != nil {
		dbSale.Author = *sale.Author
	}
	if sale.Encargado != nil {
		dbSale.Encargado = *sale.Encargado
	}
	if sale.Comision != nil {
		dbSale.Comision = *sale.Comision
	}
	if sale.Tipocambio != nil {
		dbSale.Tipocambio = *sale.Tipocambio
	}
	if sale.ComisionPagada != nil {
		dbSale.ComisionPagada = *sale.ComisionPagada
	}
	if sale.CompanyId != nil {
		dbSale.CompanyId = *sale.CompanyId
	}
	err = p.repository.Update(ctx, ID, models.Sale(dbSale))

	return err

}

// Delete user
func (p *saleService) Delete(ctx context.Context, ID string) (err error) {
	err = p.repository.Delete(ctx, ID)
	if errors.Is(err, wrappers.NonExistentErr) {
		err = wrappers.NewNonExistentErr(fmt.Errorf("ID %s not found", ID))
	}

	return err
}
