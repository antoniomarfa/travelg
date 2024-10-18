package models

import "time"

type Users struct {
	ID          string    `json:"_id,omitempty"`
	Username    string    `json:"username"`
	Name        string    `json:"name"`
	Password    string    `json:"password"`
	Email       string    `json:"email"`
	Phone       string    `json:"phone"`
	RolesId     int64     `json:"roles_id"`
	Active      int       `json:"active"`
	Author      string    `json:"author"`
	Company_id  int64     `json:"company_id"`
	CreatedDate time.Time `gorm:"autoCreateTime"`
	UpdatedDate time.Time `gorm:"autoUpdateTime"`
}

// Resp  response struct
type UsersResp struct {
	ID          string    `json:"id"`
	Username    string    `json:"username"`
	Name        string    `json:"name"`
	Password    string    `json:"password"`
	Email       string    `json:"email"`
	Phone       string    `json:"phone"`
	RolesId     int64     `json:"roles_id"`
	Active      int       `json:"active"`
	Author      string    `json:"author"`
	Company_id  int64     `json:"company_id"`
	CreatedDate time.Time `gorm:"autoCreateTime"`
	UpdatedDate time.Time `gorm:"autoUpdateTime"`
}

// Create---Req  request struct
type CreateUsersReq struct {
	ID          string    `gorm:"primaryKey;autoIncrement"`
	Username    string    `json:"username"`
	Name        string    `json:"name"`
	Password    string    `json:"password"`
	Email       string    `json:"email"`
	Phone       string    `json:"phone"`
	RolesId     int64     `json:"roles_id"`
	Active      int       `json:"active"`
	Author      string    `json:"author"`
	Company_id  int64     `json:"company_id"`
	CreatedDate time.Time `gorm:"autoCreateTime"`
	UpdatedDate time.Time `gorm:"autoUpdateTime"`
}

type UpdateUsersReq struct {
	ID          string     `json:"-"`
	Username    *string    `json:"username"`
	Name        *string    `json:"name"`
	Password    *string    `json:"password"`
	Email       *string    `json:"email"`
	Phone       *string    `json:"phone"`
	RolesId     *int64     `json:"roles_id"`
	Active      *int       `json:"active"`
	Author      *string    `json:"author"`
	Company_id  *int64     `json:"company_id"`
	CreatedDate *time.Time `gorm:"autoCreateTime"`
	UpdatedDate *time.Time `gorm:"autoUpdateTime"`
}
