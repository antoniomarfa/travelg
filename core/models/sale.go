package models

import "time"

type Sale struct {
	ID                string    `json:"_id,omitempty"`
	Fecha             time.Time `json:"fecha"`
	SellerId          int64     `json:"seller_id"`
	Identificador     string    `json:"identificador"`
	EstablecimientoId int64     `json:"establecimiento_id"`
	ProgramId         int64     `json:"program_id"`
	Curso             int       `json:"curso"`
	Idcurso           string    `json:"idcurso"`
	Nroalumno         int       `json:"nroalumno"`
	Liberados         int       `json:"liberados"`
	Program           int       `json:"program"`
	Subtotal          int       `json:"subtotal"`
	Descm             int       `json:"descm"`
	Vprograma         int       `json:"vprograma"`
	Description       string    `json:"descrition"`
	Obs               string    `json:"obs"`
	Fechasalida       time.Time `json:"fechasalida"`
	Activo            int       `json:"activo"`
	State             string    `json:"state"`
	CorreoEncargado   string    `json:"correo_encargado"`
	Password          string    `json:"password"`
	FechaUltpag       time.Time `json:"fecha_ultpag"`
	FechaCierre       time.Time `json:"fecha_cierre"`
	Sendemail         int       `json:"sendemail"`
	Author            string    `json:"author"`
	Encargado         string    `json:"encargado"`
	Comision          float32   `json:"comsion"`
	Tipocambio        float32   `json:"tipocambio"`
	ComisionPagada    int       `json:"comision_pagada"`
	CompanyId         int64     `json:"company_id"`
	CreatedDate       time.Time `gorm:"autoCreateTime"`
	UpdatedDate       time.Time `gorm:"autoUpdateTime"`
}

// Resp  response struct
type SaleResp struct {
	ID                string    `json:"id"`
	Fecha             time.Time `json:"fecha"`
	SellerId          int64     `json:"seller_id"`
	Identificador     string    `json:"identificador"`
	EstablecimientoId int64     `json:"establecimiento_id"`
	ProgramId         int64     `json:"program_id"`
	Curso             int       `json:"curso"`
	Idcurso           string    `json:"idcurso"`
	Nroalumno         int       `json:"nroalumno"`
	Liberados         int       `json:"liberados"`
	Program           int       `json:"program"`
	Subtotal          int       `json:"subtotal"`
	Descm             int       `json:"descm"`
	Vprograma         int       `json:"vprograma"`
	Description       string    `json:"descrition"`
	Obs               string    `json:"obs"`
	Fechasalida       time.Time `json:"fechasalida"`
	Activo            int       `json:"activo"`
	State             string    `json:"state"`
	CorreoEncargado   string    `json:"correo_encargado"`
	Password          string    `json:"password"`
	FechaUltpag       time.Time `json:"fecha_ultpag"`
	FechaCierre       time.Time `json:"fecha_cierre"`
	Sendemail         int       `json:"sendemail"`
	Author            string    `json:"author"`
	Encargado         string    `json:"encargado"`
	Comision          float32   `json:"comsion"`
	Tipocambio        float32   `json:"tipocambio"`
	ComisionPagada    int       `json:"comision_pagada"`
	CompanyId         int64     `json:"company_id"`
	CreatedDate       time.Time `gorm:"autoCreateTime"`
	UpdatedDate       time.Time `gorm:"autoUpdateTime"`
}

// Create---Req  request struct
type CreateSaleReq struct {
	ID                string    `gorm:"primaryKey;autoIncrement"`
	Fecha             time.Time `json:"fecha"`
	SellerId          int64     `json:"seller_id"`
	Identificador     string    `json:"identificador"`
	EstablecimientoId int64     `json:"establecimiento_id"`
	ProgramId         int64     `json:"program_id"`
	Curso             int       `json:"curso"`
	Idcurso           string    `json:"idcurso"`
	Nroalumno         int       `json:"nroalumno"`
	Liberados         int       `json:"liberados"`
	Program           int       `json:"program"`
	Subtotal          int       `json:"subtotal"`
	Descm             int       `json:"descm"`
	Vprograma         int       `json:"vprograma"`
	Description       string    `json:"descrition"`
	Obs               string    `json:"obs"`
	Fechasalida       time.Time `json:"fechasalida"`
	Activo            int       `json:"activo"`
	State             string    `json:"state"`
	CorreoEncargado   string    `json:"correo_encargado"`
	Password          string    `json:"password"`
	FechaUltpag       time.Time `json:"fecha_ultpag"`
	FechaCierre       time.Time `json:"fecha_cierre"`
	Sendemail         int       `json:"sendemail"`
	Author            string    `json:"author"`
	Encargado         string    `json:"encargado"`
	Comision          float32   `json:"comsion"`
	Tipocambio        float32   `json:"tipocambio"`
	ComisionPagada    int       `json:"comision_pagada"`
	CompanyId         int64     `json:"company_id"`
	CreatedDate       time.Time `gorm:"autoCreateTime"`
	UpdatedDate       time.Time `gorm:"autoUpdateTime"`
}

type UpdateSaleReq struct {
	ID                string     `json:"-"`
	Fecha             *time.Time `json:"fecha"`
	SellerId          *int64     `json:"seller_id"`
	Identificador     *string    `json:"identificador"`
	EstablecimientoId *int64     `json:"establecimiento_id"`
	ProgramId         *int64     `json:"program_id"`
	Curso             *int       `json:"curso"`
	Idcurso           *string    `json:"idcurso"`
	Nroalumno         *int       `json:"nroalumno"`
	Liberados         *int       `json:"liberados"`
	Program           *int       `json:"program"`
	Subtotal          *int       `json:"subtotal"`
	Descm             *int       `json:"descm"`
	Vprograma         *int       `json:"vprograma"`
	Description       *string    `json:"descrition"`
	Obs               *string    `json:"obs"`
	Fechasalida       *time.Time `json:"fechasalida"`
	Activo            *int       `json:"activo"`
	State             *string    `json:"state"`
	CorreoEncargado   *string    `json:"correo_encargado"`
	Password          *string    `json:"password"`
	FechaUltpag       *time.Time `json:"fecha_ultpag"`
	FechaCierre       *time.Time `json:"fecha_cierre"`
	Sendemail         *int       `json:"sendemail"`
	Author            *string    `json:"author"`
	Encargado         *string    `json:"encargado"`
	Comision          *float32   `json:"comsion"`
	Tipocambio        *float32   `json:"tipocambio"`
	ComisionPagada    *int       `json:"comision_pagada"`
	CompanyId         *int64     `json:"company_id"`
	CreatedDate       *time.Time `gorm:"autoCreateTime"`
	UpdatedDate       *time.Time `gorm:"autoUpdateTime"`
}
