package models

import (
	"time"
)

type Company struct {
	ID              string    `json:"_id,omitempty"`
	Rut             string    `json:"Rut"`
	Razonsocial     string    `json:"Razonsocial"`
	Nomfantasia     string    `json:"Nomfantasia"`
	Rutreplegal     string    `json:"Rutreplegal"`
	Replegal        string    `json:"Replegal"`
	Contrato        string    `json:"Contrato"`
	ActiveFlow      int       `json:"Active_flow"`
	FlowApikey      string    `json:"flow_apikey"`
	FlowSecretkey   string    `json:"flow_secretkey"`
	ActiveTrb       int       `json:"active_trb"`
	TrbCommercecode string    `json:"trb_commercecode"`
	ComunaId        int64     `json:"comuna_id"`
	RegionId        int64     `json:"region_id"`
	Fono            string    `json:"Fono"`
	Correo          string    `json:"Correo"`
	ContactoNombre1 string    `json:"contacto_nombre1"`
	ContactoFono1   string    `json:"contacto_fono1"`
	ContactoCorreo1 string    `json:"contacto_correo1"`
	ContactoNombre2 string    `json:"contacto_nombre2"`
	ContactoFono2   string    `json:"contacto_fono2"`
	ContactoCorreo2 string    `json:"contacto_correo2"`
	Author          string    `json:"author"`
	Active          string    `json:"active"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

// Resp  response struct
type CompanyResp struct {
	ID              string    `json:"id"`
	Rut             string    `json:"Rut"`
	Razonsocial     string    `json:"Razonsocial"`
	Nomfantasia     string    `json:"Nomfantasia"`
	Rutreplegal     string    `json:"Rutreplegal"`
	Replegal        string    `json:"Replegal"`
	Contrato        string    `json:"Contrato"`
	ActiveFlow      int       `json:"Active_flow"`
	FlowApikey      string    `json:"flow_apikey"`
	FlowSecretkey   string    `json:"flow_secretkey"`
	ActiveTrb       int       `json:"active_trb"`
	TrbCommercecode string    `json:"trb_commercecode"`
	ComunaId        int64     `json:"comuna_id"`
	RegionId        int64     `json:"region_id"`
	Fono            string    `json:"Fono"`
	Correo          string    `json:"Correo"`
	ContactoNombre1 string    `json:"contacto_nombre1"`
	ContactoFono1   string    `json:"contacto_fono1"`
	ContactoCorreo1 string    `json:"contacto_correo1"`
	ContactoNombre2 string    `json:"contacto_nombre2"`
	ContactoFono2   string    `json:"contacto_fono2"`
	ContactoCorreo2 string    `json:"contacto_correo2"`
	Author          string    `json:"author"`
	Active          string    `json:"active"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

// Create---Req  request struct
type CreateCompanyReq struct {
	ID              string    `gorm:"primaryKey;autoIncrement"`
	Rut             string    `json:"Rut"`
	Razonsocial     string    `json:"Razonsocial"`
	Nomfantasia     string    `json:"Nomfantasia"`
	Rutreplegal     string    `json:"Rutreplegal"`
	Replegal        string    `json:"Replegal"`
	Contrato        string    `json:"Contrato"`
	ActiveFlow      int       `json:"Active_flow"`
	FlowApikey      string    `json:"flow_apikey"`
	FlowSecretkey   string    `json:"flow_secretkey"`
	ActiveTrb       int       `json:"active_trb"`
	TrbCommercecode string    `json:"trb_commercecode"`
	ComunaId        int64     `json:"comuna_id"`
	RegionId        int64     `json:"region_id"`
	Fono            string    `json:"Fono"`
	Correo          string    `json:"Correo"`
	ContactoNombre1 string    `json:"contacto_nombre1"`
	ContactoFono1   string    `json:"contacto_fono1"`
	ContactoCorreo1 string    `json:"contacto_correo1"`
	ContactoNombre2 string    `json:"contacto_nombre2"`
	ContactoFono2   string    `json:"contacto_fono2"`
	ContactoCorreo2 string    `json:"contacto_correo2"`
	Author          string    `json:"author"`
	Active          string    `json:"active"`
	CreatedAt       time.Time `gorm:"autoCreateTime"`
	UpdatedAt       time.Time `gorm:"autoUpdateTime"`
}

type UpdateCompanyReq struct {
	ID              string    `json:"-"`
	Rut             *string   `json:"Rut"`
	Razonsocial     *string   `json:"Razonsocial"`
	Nomfantasia     *string   `json:"Nomfantasia"`
	Rutreplegal     *string   `json:"Rutreplegal"`
	Replegal        *string   `json:"Replegal"`
	Contrato        *string   `json:"Contrato"`
	ActiveFlow      *int      `json:"Active_flow"`
	FlowApikey      *string   `json:"flow_apikey"`
	FlowSecretkey   *string   `json:"flow_secretkey"`
	ActiveTrb       *int      `json:"active_trb"`
	TrbCommercecode *string   `json:"trb_commercecode"`
	ComunaId        *int64    `json:"comuna_id"`
	RegionId        *int64    `json:"region_id"`
	Fono            *string   `json:"Fono"`
	Correo          *string   `json:"Correo"`
	ContactoNombre1 *string   `json:"contacto_nombre1"`
	ContactoFono1   *string   `json:"contacto_fono1"`
	ContactoCorreo1 *string   `json:"contacto_correo1"`
	ContactoNombre2 *string   `json:"contacto_nombre2"`
	ContactoFono2   *string   `json:"contacto_fono2"`
	ContactoCorreo2 *string   `json:"contacto_correo2"`
	Author          *string   `json:"author"`
	Active          *string   `json:"active"`
	CreatedAt       time.Time `gorm:"autoCreateTime"`
	UpdatedAt       time.Time `gorm:"autoUpdateTime"`
}
