package models

import "time"

type Curso struct {
	ID             string    `json:"_id,omitempty"`
	SaleId         int8      `json:"sale_id"`
	Rutalumno      string    `json:"rutalumno"`
	Nombrealumno   string    `json:"nombrealumno"`
	Fechanac       time.Time `json:"fechanac"`
	Rutapod        string    `json:"rutapod"`
	Nombreapod     string    `json:"nombreapod"`
	Dircalle       string    `json:"dircalle"`
	Dirnumero      string    `json:"dirnumero"`
	Nrodepto       string    `json:"nrodepto"`
	Region_id      int64     `json:"region_id"`
	Comuna_id      int64     `json:"comuna_id"`
	Fono           string    `json:"fono"`
	Celular        string    `json:"celular"`
	Correo         string    `json:"correo"`
	Vpagar         int       `json:"vpagar"`
	Descto         int       `json:"descto"`
	Apagar         int       `json:"apagar"`
	Liberado       int       `json:"liberado"`
	Enviado        string    `json:"enviado"`
	Estado         string    `json:"estado"`
	Password       string    `json:"password"`
	AceptaContrato int       `json:"acepta_contrato"`
	Signature      string    `json:"signature"`
	Author         string    `json:"author"`
	CompanyId      int64     `json:"company_id"`
	CreatedDate    time.Time `json:"created_date"`
	UpdatedDate    time.Time `json:"updated_date"`
}

// Resp  response struct
type CursoResp struct {
	ID             string    `json:"id"`
	SaleId         int8      `json:"sale_id"`
	Rutalumno      string    `json:"rutalumno"`
	Nombrealumno   string    `json:"nombrealumno"`
	Fechanac       time.Time `json:"fechanac"`
	Rutapod        string    `json:"rutapod"`
	Nombreapod     string    `json:"nombreapod"`
	Dircalle       string    `json:"dircalle"`
	Dirnumero      string    `json:"dirnumero"`
	Nrodepto       string    `json:"nrodepto"`
	Region_id      int64     `json:"region_id"`
	Comuna_id      int64     `json:"comuna_id"`
	Fono           string    `json:"fono"`
	Celular        string    `json:"celular"`
	Correo         string    `json:"correo"`
	Vpagar         int       `json:"vpagar"`
	Descto         int       `json:"descto"`
	Apagar         int       `json:"apagar"`
	Liberado       int       `json:"liberado"`
	Enviado        string    `json:"enviado"`
	Estado         string    `json:"estado"`
	Password       string    `json:"password"`
	AceptaContrato int       `json:"acepta_contrato"`
	Signature      string    `json:"signature"`
	Author         string    `json:"author"`
	CompanyId      int64     `json:"company_id"`
	CreatedDate    time.Time `json:"created_date"`
	UpdatedDate    time.Time `json:"updated_date"`
}

// Create---Req  request struct
type CreateCursoReq struct {
	ID             string    `gorm:"primaryKey;autoIncrement"`
	SaleId         int8      `json:"sale_id"`
	Rutalumno      string    `json:"rutalumno"`
	Nombrealumno   string    `json:"nombrealumno"`
	Fechanac       time.Time `json:"fechanac"`
	Rutapod        string    `json:"rutapod"`
	Nombreapod     string    `json:"nombreapod"`
	Dircalle       string    `json:"dircalle"`
	Dirnumero      string    `json:"dirnumero"`
	Nrodepto       string    `json:"nrodepto"`
	Region_id      int64     `json:"region_id"`
	Comuna_id      int64     `json:"comuna_id"`
	Fono           string    `json:"fono"`
	Celular        string    `json:"celular"`
	Correo         string    `json:"correo"`
	Vpagar         int       `json:"vpagar"`
	Descto         int       `json:"descto"`
	Apagar         int       `json:"apagar"`
	Liberado       int       `json:"liberado"`
	Enviado        string    `json:"enviado"`
	Estado         string    `json:"estado"`
	Password       string    `json:"password"`
	AceptaContrato int       `json:"acepta_contrato"`
	Signature      string    `json:"signature"`
	Author         string    `json:"author"`
	CompanyId      int64     `json:"company_id"`
	CreatedDate    time.Time `gorm:"autoCreateTime"`
	UpdatedDate    time.Time `gorm:"autoUpdateTime"`
}

type UpdateCursoReq struct {
	ID             string     `json:"-"`
	SaleId         *int8      `json:"sale_id"`
	Rutalumno      *string    `json:"rutalumno"`
	Nombrealumno   *string    `json:"nombrealumno"`
	Fechanac       *time.Time `json:"fechanac"`
	Rutapod        *string    `json:"rutapod"`
	Nombreapod     *string    `json:"nombreapod"`
	Dircalle       *string    `json:"dircalle"`
	Dirnumero      *string    `json:"dirnumero"`
	Nrodepto       *string    `json:"nrodepto"`
	Region_id      *int64     `json:"region_id"`
	Comuna_id      *int64     `json:"comuna_id"`
	Fono           *string    `json:"fono"`
	Celular        *string    `json:"celular"`
	Correo         *string    `json:"correo"`
	Vpagar         *int       `json:"vpagar"`
	Descto         *int       `json:"descto"`
	Apagar         *int       `json:"apagar"`
	Liberado       *int       `json:"liberado"`
	Enviado        *string    `json:"enviado"`
	Estado         *string    `json:"estado"`
	Password       *string    `json:"password"`
	AceptaContrato *int       `json:"acepta_contrato"`
	Signature      *string    `json:"signature"`
	Author         *string    `json:"author"`
	CompanyId      *int64     `json:"company_id"`
	CreatedDate    *time.Time `gorm:"autoCreateTime"`
	UpdatedDate    *time.Time `gorm:"autoUpdateTime"`
}
