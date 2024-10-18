package models

import "time"

type Ingreso struct {
	ID            string    `json:"_id,omitempty"`
	Tipocomp      string    `json:"tipocomp"`
	Fecha         time.Time `json:"fecha"`
	Identificador string    `json:"identificador"`
	SaleId        int64     `json:"sale_id"`
	CursoId       int64     `json:"curso_id"`
	Rutapo        string    `json:"rutapo"`
	Rutalum       string    `json:"rutalum"`
	Fpago         string    `json:"fpago"`
	Monto         int       `json:"monto"`
	Activo        int       `json:"activo"`
	StatusPago    string    `json:"status_pago"`
	Author        string    `json:"author"`
	CompanyId     int64     `json:"company_id"`
	CreatedDate   time.Time `gorm:"autoCreateTime"`
	UpdatedDate   time.Time `gorm:"autoUpdateTime"`
}

// Resp  response struct
type IngresoResp struct {
	ID            string    `json:"id"`
	Tipocomp      string    `json:"tipocomp"`
	Fecha         time.Time `json:"fecha"`
	Identificador string    `json:"identificador"`
	SaleId        int64     `json:"sale_id"`
	CursoId       int64     `json:"curso_id"`
	Rutapo        string    `json:"rutapo"`
	Rutalum       string    `json:"rutalum"`
	Fpago         string    `json:"fpago"`
	Monto         int       `json:"monto"`
	Activo        int       `json:"activo"`
	StatusPago    string    `json:"status_pago"`
	Author        string    `json:"author"`
	CompanyId     int64     `json:"company_id"`
	CreatedDate   time.Time `gorm:"autoCreateTime"`
	UpdatedDate   time.Time `gorm:"autoUpdateTime"`
}

// Create---Req  request struct
type CreateIngresoReq struct {
	ID            string    `gorm:"primaryKey;autoIncrement"`
	Tipocomp      string    `json:"tipocomp"`
	Fecha         time.Time `json:"fecha"`
	Identificador string    `json:"identificador"`
	SaleId        int64     `json:"sale_id"`
	CursoId       int64     `json:"curso_id"`
	Rutapo        string    `json:"rutapo"`
	Rutalum       string    `json:"rutalum"`
	Fpago         string    `json:"fpago"`
	Monto         int       `json:"monto"`
	Activo        int       `json:"activo"`
	StatusPago    string    `json:"status_pago"`
	Author        string    `json:"author"`
	CompanyId     int64     `json:"company_id"`
	CreatedDate   time.Time `gorm:"autoCreateTime"`
	UpdatedDate   time.Time `gorm:"autoUpdateTime"`
}

type UpdateIngresoReq struct {
	ID            string     `json:"-"`
	Tipocomp      *string    `json:"tipocomp"`
	Fecha         *time.Time `json:"fecha"`
	Identificador *string    `json:"identificador"`
	SaleId        *int64     `json:"sale_id"`
	CursoId       *int64     `json:"curso_id"`
	Rutapo        *string    `json:"rutapo"`
	Rutalum       *string    `json:"rutalum"`
	Fpago         *string    `json:"fpago"`
	Monto         *int       `json:"monto"`
	Activo        *int       `json:"activo"`
	StatusPago    *string    `json:"status_pago"`
	Author        *string    `json:"author"`
	CompanyId     *int64     `json:"company_id"`
	CreatedDate   *time.Time `gorm:"autoCreateTime"`
	UpdatedDate   *time.Time `gorm:"autoUpdateTime"`
}
