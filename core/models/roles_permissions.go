package models

type RolesPermissions struct {
	ID         string `json:"_id,omitempty"`
	RolesId    int64  `json:"roles_id"`
	Permission string `json:"permission"`
	Actions    string `json:"actions"`
	CompanyId  int64  `json:"company_id"`
}

// Resp  response struct
type RolesPermissionsResp struct {
	ID         string `json:"id"`
	RolesId    int64  `json:"roles_id"`
	Permission string `json:"permission"`
	Actions    string `json:"actions"`
	CompanyId  int64  `json:"company_id"`
}

// Create---Req  request struct
type CreateRolesPermissionsReq struct {
	ID         string `gorm:"primaryKey;autoIncrement"`
	RolesId    int64  `json:"roles_id"`
	Permission string `json:"permission"`
	Actions    string `json:"actions"`
	CompanyId  int64  `json:"company_id"`
}

type UpdateRolesPermissionsReq struct {
	ID         string  `json:"-"`
	RolesId    *int64  `json:"roles_id"`
	Permission *string `json:"permission"`
	Actions    *string `json:"actions"`
	CompanyId  *int64  `json:"company_id"`
}
