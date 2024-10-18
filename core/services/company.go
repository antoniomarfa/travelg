package services

import (
	"context"
	"errors"
	"fmt"
	"time"

	"travel/config"
	"travel/core/models"
	"travel/core/ports"
	"travel/tools/wrappers"
)

// rolesService adapter of an user service
type companyService struct {
	config     config.Config
	repository ports.CompanyRepository
}

// NewURolesService creates a new user service
func NewCompanyService(cfg config.Config, repo ports.CompanyRepository) ports.CompanyService {
	return &companyService{
		config:     cfg,
		repository: repo,
	}
}

// Create roles
func (p *companyService) Create(ctx context.Context, company models.CreateCompanyReq) (resp models.CreationResp, err error) {

	now := time.Now().UTC()
	company.CreatedAt = now
	company.UpdatedAt = now
	insertedID, err := p.repository.Create(ctx, models.CreateCompanyReq(company))
	if err != nil {
		return
	}

	resp = models.CreationResp{
		InsertedID: insertedID,
	}

	return
}

// GetAll users
func (p *companyService) GetAll(ctx context.Context) (resp []models.CompanyResp, err error) {
	// Obtiene los roles desde el repositorio
	result, err := p.repository.Get(ctx, map[string]interface{}{}, nil, nil)
	if err != nil {
		return nil, err
	}

	// Convierte los resultados a la estructura de respuesta
	resp = make([]models.CompanyResp, len(result))
	for i, v := range result {
		// Asegúrate de que el tipo sea correcto antes de hacer la conversión
		if company, ok := v.(models.CompanyResp); ok {
			resp[i] = models.CompanyResp(company) // Asumiendo que tienes un constructor o un mapeo en tu struct
		} else {
			return nil, wrappers.NewNonExistentErr(fmt.Errorf("error de tipo: se esperaba entities.Products, pero se obtuvo %T", v))
		}
	}
	return
}

// GetByID user
func (p *companyService) GetByID(ctx context.Context, ID string) (resp models.CompanyResp, err error) {
	company, err := p.repository.GetByID(ctx, ID)

	if err != nil {
		return resp, fmt.Errorf("colegio con ID %s no encontrado", ID)
	}

	if company == nil {
		// Si no se encuentra el colegio (colegios es nil), devolver un valor en blanco y un error
		return models.CompanyResp{}, fmt.Errorf("colegio con ID %s no encontrado", ID)
	}

	resp = *company.(*models.CompanyResp)

	return
}

// Update user
func (p *companyService) Update(ctx context.Context, ID string, company models.UpdateCompanyReq) (err error) {

	dbCompany, err := p.GetByID(ctx, ID)
	if err != nil {
		return
	}
	// Actualizar los campos solo si no son nil
	if company.Rut != nil {
		dbCompany.Rut = *company.Rut
	}

	if company.Razonsocial != nil {
		dbCompany.Razonsocial = *company.Razonsocial
	}

	if company.Nomfantasia != nil {
		dbCompany.Nomfantasia = *company.Nomfantasia
	}

	if company.Rutreplegal != nil {
		dbCompany.Rutreplegal = *company.Rutreplegal
	}

	if company.Replegal != nil {
		dbCompany.Replegal = *company.Replegal
	}

	if company.Contrato != nil {
		dbCompany.Contrato = *company.Contrato
	}

	if company.ActiveFlow != nil {
		dbCompany.ActiveFlow = *company.ActiveFlow
	}

	if company.FlowApikey != nil {
		dbCompany.FlowApikey = *company.FlowApikey
	}

	if company.FlowSecretkey != nil {
		dbCompany.FlowSecretkey = *company.FlowSecretkey
	}

	if company.ActiveTrb != nil {
		dbCompany.ActiveTrb = *company.ActiveTrb
	}

	if company.TrbCommercecode != nil {
		dbCompany.TrbCommercecode = *company.TrbCommercecode
	}

	if company.ComunaId != nil {
		dbCompany.ComunaId = *company.ComunaId
	}

	if company.RegionId != nil {
		dbCompany.RegionId = *company.RegionId
	}

	if company.Fono != nil {
		dbCompany.Fono = *company.Fono
	}

	if company.Correo != nil {
		dbCompany.Correo = *company.Correo
	}

	if company.ContactoNombre1 != nil {
		dbCompany.ContactoNombre1 = *company.ContactoNombre1
	}

	if company.ContactoFono1 != nil {
		dbCompany.ContactoFono1 = *company.ContactoFono1
	}

	if company.ContactoCorreo1 != nil {
		dbCompany.ContactoCorreo1 = *company.ContactoCorreo1
	}

	if company.ContactoNombre2 != nil {
		dbCompany.ContactoNombre2 = *company.ContactoNombre2
	}

	if company.ContactoFono2 != nil {
		dbCompany.ContactoFono2 = *company.ContactoFono2
	}

	if company.ContactoCorreo2 != nil {
		dbCompany.ContactoCorreo2 = *company.ContactoCorreo2
	}

	if company.Author != nil {
		dbCompany.Author = *company.Author
	}

	// Asegúrate de que Active no sea nil antes de asignarlo
	if company.Active != nil {
		dbCompany.Active = *company.Active
	}

	// Actualizar la fecha de modificación
	dbCompany.UpdatedAt = time.Now().UTC()

	// Llamar al repositorio para actualizar la entidad
	err = p.repository.Update(ctx, ID, models.Company(dbCompany))

	return err
}

// Delete user
func (p *companyService) Delete(ctx context.Context, ID string) (err error) {
	err = p.repository.Delete(ctx, ID)
	if errors.Is(err, wrappers.NonExistentErr) {
		err = wrappers.NewNonExistentErr(fmt.Errorf("ID %s not found", ID))
	}

	return err
}
