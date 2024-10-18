package models

import "time"

type Pagos struct {
	ID            string      `json:"_id,omitempty"`
	Tipocom       string      `json:"tipocom"`
	IngresoId     int64       `json:"ingreso_id"`
	Identificador string      `json:"identificador"`
	Fecha         time.Time   `json:"fecha"`
	SaleId        int64       `json:"sale_id"`
	Rutalumn      string      `json:"rutalumn"`
	Transaccion   string      `json:"transaccion"`
	Tipo          string      `json:"tipo"`
	Monto         int         `json:"monto"`
	Nrotarjeta    string      `json:"nrotarjeta"`
	Codigoauto    string      `json:"codigoauto"`
	Fechaauto     time.Ticker `json:"fechaauto"`
	Tipopago      string      `json:"tipopago"`
	Nrocuota      int         `json:"nrocuota"`
	Fechatransac  time.Time   `json:"fechatransac"`
	Author        string      `json:"author"`
	Activo        int         `json:"activo"`
	CompanyId     int64       `json:"company_id"`
	CreatedDate   time.Time   `gorm:"autoCreateTime"`
	UpdatedDate   time.Time   `gorm:"autoUpdateTime"`
}

// Resp  response struct
type PagosResp struct {
	ID            string      `json:"id"`
	Tipocom       string      `json:"tipocom"`
	IngresoId     int64       `json:"ingreso_id"`
	Identificador string      `json:"identificador"`
	Fecha         time.Time   `json:"fecha"`
	SaleId        int64       `json:"sale_id"`
	Rutalumn      string      `json:"rutalumn"`
	Transaccion   string      `json:"transaccion"`
	Tipo          string      `json:"tipo"`
	Monto         int         `json:"monto"`
	Nrotarjeta    string      `json:"nrotarjeta"`
	Codigoauto    string      `json:"codigoauto"`
	Fechaauto     time.Ticker `json:"fechaauto"`
	Tipopago      string      `json:"tipopago"`
	Nrocuota      int         `json:"nrocuota"`
	Fechatransac  time.Time   `json:"fechatransac"`
	Author        string      `json:"author"`
	Activo        int         `json:"activo"`
	CompanyId     int64       `json:"company_id"`
	CreatedDate   time.Time   `gorm:"autoCreateTime"`
	UpdatedDate   time.Time   `gorm:"autoUpdateTime"`
}

// Create---Req  request struct
type CreatePagosReq struct {
	ID            string      `gorm:"primaryKey;autoIncrement"`
	Tipocom       string      `json:"tipocom"`
	IngresoId     int64       `json:"ingreso_id"`
	Identificador string      `json:"identificador"`
	Fecha         time.Time   `json:"fecha"`
	SaleId        int64       `json:"sale_id"`
	Rutalumn      string      `json:"rutalumn"`
	Transaccion   string      `json:"transaccion"`
	Tipo          string      `json:"tipo"`
	Monto         int         `json:"monto"`
	Nrotarjeta    string      `json:"nrotarjeta"`
	Codigoauto    string      `json:"codigoauto"`
	Fechaauto     time.Ticker `json:"fechaauto"`
	Tipopago      string      `json:"tipopago"`
	Nrocuota      int         `json:"nrocuota"`
	Fechatransac  time.Time   `json:"fechatransac"`
	Author        string      `json:"author"`
	Activo        int         `json:"activo"`
	CompanyId     int64       `json:"company_id"`
	CreatedDate   time.Time   `gorm:"autoCreateTime"`
	UpdatedDate   time.Time   `gorm:"autoUpdateTime"`
}

type UpdatePagosReq struct {
	ID            string       `json:"-"`
	Tipocom       *string      `json:"tipocom"`
	IngresoId     *int64       `json:"ingreso_id"`
	Identificador *string      `json:"identificador"`
	Fecha         *time.Time   `json:"fecha"`
	SaleId        *int64       `json:"sale_id"`
	Rutalumn      *string      `json:"rutalumn"`
	Transaccion   *string      `json:"transaccion"`
	Tipo          *string      `json:"tipo"`
	Monto         *int         `json:"monto"`
	Nrotarjeta    *string      `json:"nrotarjeta"`
	Codigoauto    *string      `json:"codigoauto"`
	Fechaauto     *time.Ticker `json:"fechaauto"`
	Tipopago      *string      `json:"tipopago"`
	Nrocuota      *int         `json:"nrocuota"`
	Fechatransac  *time.Time   `json:"fechatransac"`
	Author        *string      `json:"author"`
	Activo        *int         `json:"activo"`
	CompanyId     *int64       `json:"company_id"`
	CreatedDate   *time.Time   `gorm:"autoCreateTime"`
	UpdatedDate   *time.Time   `gorm:"autoUpdateTime"`
}
