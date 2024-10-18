package models

import "time"

type Roles struct {
	ID          string    `json:"_id,omitempty"`
	Description string    `json:"description"`
	Active      int       `json:"active"`
	CreatedDate time.Time `gorm:"autoCreateTime"`
	UpdatedDate time.Time `gorm:"autoUpdateTime"`
	Author      string    `json:"author"`
	CompanyId   int8      `json:"company_id"`
}

// Resp  response struct
type RolesResp struct {
	ID          string    `json:"id"`
	Description string    `json:"description"`
	Active      int       `json:"active"`
	CreatedDate time.Time `gorm:"autoCreateTime"`
	UpdatedDate time.Time `gorm:"autoUpdateTime"`
	Author      string    `json:"author"`
	CompanyId   int8      `json:"company_id"`
}

// Create---Req  request struct
type CreateRolesReq struct {
	ID          string    `gorm:"primaryKey;autoIncrement"`
	Description string    `json:"description"`
	Active      int       `json:"active"`
	CreatedDate time.Time `gorm:"autoCreateTime"`
	UpdatedDate time.Time `gorm:"autoUpdateTime"`
	Author      string    `json:"author"`
	CompanyId   int8      `json:"company_id"`
}

type UpdateRolesReq struct {
	ID          string     `json:"-"`
	Description *string    `json:"description"`
	Active      *int       `json:"active"`
	CreatedDate *time.Time `gorm:"autoCreateTime"`
	UpdatedDate *time.Time `gorm:"autoUpdateTime"`
	Author      *string    `json:"author"`
	CompanyId   *int8      `json:"company_id"`
}
