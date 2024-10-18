package entities

import "time"

// EntityNameRoles contains the name of the entity
const EntityNameCompany = "company"

type Company struct {
	ID              string    `bson:"_id,omitempty"`
	Rut             string    `bson:"Rut"`
	Razonsocial     string    `bson:"Razonsocial"`
	Nomfantasia     string    `bson:"Nomfantasia"`
	Rutreplegal     string    `bson:"Rutreplegal"`
	Replegal        string    `bson:"Replegal"`
	Contrato        string    `bson:"Contrato"`
	ActiveFlow      int       `bson:"Active_flow"`
	FlowApikey      string    `bson:"flow_apikey"`
	FlowSecretkey   string    `bson:"flow_secretkey"`
	ActiveTrb       int       `bson:"active_trb"`
	TrbCommercecode string    `bson:"trb_commercecode"`
	ComunaId        int64     `bson:"comuna_id"`
	RegionId        int64     `bson:"region_id"`
	Fono            string    `bson:"Fono"`
	Correo          string    `bson:"Correo"`
	ContactoNombre1 string    `bson:"contacto_nombre1"`
	ContactoFono1   string    `bson:"contacto_fono1"`
	ContactoCorreo1 string    `bson:"contacto_correo1"`
	ContactoNombre2 string    `bson:"contacto_nombre2"`
	ContactoFono2   string    `bson:"contacto_fono2"`
	ContactoCorreo2 string    `bson:"contacto_correo2"`
	Author          string    `bson:"author"`
	Active          string    `bson:"active"`
	CreatedAt       time.Time `bson:"created_at"`
	UpdatedAt       time.Time `bson:"updated_at"`
}
