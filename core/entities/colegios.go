package entities

// EntityNameRoles contains the name of the entity
const EntityNameColegio = "colegios"

type Colegios struct {
	ID        string `bson:"id"`
	Codigo    string `bson:"codigo"`
	Nombre    string `bson:"nombre"`
	Direccion string `bson:"direccion"`
	Comuna    string `bson:"comuna"`
	Latitud   int16  `bson:"latitiud"`
	Longitud  int16  `bson:"longitud"`
	RegionId  int64  `bson:"region_id"`
	ComunaId  int64  `bson:"comuna_id"`
	CompanyId int64  `bson:"company_id"`
}
