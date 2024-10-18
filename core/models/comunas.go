package models

import "time"

type Comunas struct {
	ID          string    `json:"_id,omitempty"`
	RegionsId   int8      `json:"regions_id"`
	Code        string    `json:"code"`
	Description string    `json:"description"`
	Active      int       `json:"active"`
	CreatedDate time.Time `gorm:"autoCreateTime"`
	UpdatedDate time.Time `gorm:"autoUpdateTime"`
	Author      string    `json:"author"`
}

// Resp  response struct
type ComunasResp struct {
	ID          string    `json:"id"`
	RegionsId   int8      `json:"regions_id"`
	Code        string    `json:"code"`
	Description string    `json:"description"`
	Active      int       `json:"active"`
	CreatedDate time.Time `gorm:"autoCreateTime"`
	UpdatedDate time.Time `gorm:"autoUpdateTime"`
	Author      string    `json:"author"`
}

// Create---Req  request struct
type CreateComunasReq struct {
	ID          string    `gorm:"primaryKey;autoIncrement"`
	RegionsId   int8      `json:"regions_id"`
	Code        string    `json:"code"`
	Description string    `json:"description"`
	Active      int       `json:"active"`
	CreatedDate time.Time `gorm:"autoCreateTime"`
	UpdatedDate time.Time `gorm:"autoUpdateTime"`
	Author      string    `json:"author"`
}

type UpdateComunasReq struct {
	ID          string     `json:"-"`
	RegionsId   *int8      `json:"regions_id"`
	Code        *string    `json:"code"`
	Description *string    `json:"description"`
	Active      *int       `json:"active"`
	CreatedDate *time.Time `gorm:"autoCreateTime"`
	UpdatedDate *time.Time `gorm:"autoUpdateTime"`
	Author      *string    `json:"author"`
}
