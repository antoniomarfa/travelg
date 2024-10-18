package models

import "github.com/shopspring/decimal"

type Colegios struct {
	ID        string `json:"_id,omitempty"`
	Codigo    string `json:"codigo"`
	Nombre    string `json:"nombre"`
	Direccion string `json:"direccion"`
	Comuna    string `json:"comuna"`
	Latitud   int16  `json:"latitiud"`
	Longitud  int16  `json:"longitud"`
	RegionId  int64  `json:"region_id"`
	ComunaId  int64  `json:"comuna_id"`
	CompanyId int64  `json:"company_id"`
}

// Resp  response struct
type ColegiosResp struct {
	ID        string `json:"id"`
	Codigo    string `json:"codigo"`
	Nombre    string `json:"nombre"`
	Direccion string `json:"direccion"`
	Comuna    string `json:"comuna"`
	Latitud   int16  `json:"latitiud"`
	Longitud  int16  `json:"longitud"`
	RegionId  int64  `json:"region_id"`
	ComunaId  int64  `json:"comuna_id"`
	CompanyId int64  `json:"company_id"`
}

// Create---Req  request struct
type CreateColegiosReq struct {
	ID        string          `gorm:"primaryKey;autoIncrement"`
	Codigo    string          `json:"codigo"`
	Nombre    string          `json:"nombre"`
	Direccion string          `json:"direccion"`
	Comuna    string          `json:"comuna"`
	Latitud   decimal.Decimal `json:"latitiud"`
	Longitud  decimal.Decimal `json:"longitud"`
	RegionId  int64           `json:"region_id"`
	ComunaId  int64           `json:"comuna_id"`
	CompanyId int64           `json:"company_id"`
}

type UpdateColegiosReq struct {
	ID        string  `json:"-"`
	Codigo    *string `json:"codigo"`
	Nombre    *string `json:"nombre"`
	Direccion *string `json:"direccion"`
	Comuna    *string `json:"comuna"`
	Latitud   *int16  `json:"latitiud"`
	Longitud  *int16  `json:"longitud"`
	RegionId  *int64  `json:"region_id"`
	ComunaId  *int64  `json:"comuna_id"`
	CompanyId *int64  `json:"company_id"`
}
