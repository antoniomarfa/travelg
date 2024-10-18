package models

import "time"

type Program struct {
	ID          string     `json:"_id,omitempty"`
	Code        string     `json:"code"`
	Name        string     `json:"name"`
	Desde1      int        `json:"desde1"`
	Hasta1      int        `json:"hasta1"`
	Valor1      int        `json:"valor1"`
	Liberado1   int        `json:"liberado1"`
	Desde2      int        `json:"desde2"`
	Hasta2      int        `json:"hasta2"`
	Valor2      int        `json:"valor2"`
	Liberado2   int        `json:"liberado2"`
	Desde3      int        `json:"desde3"`
	Hasta3      int        `json:"hasta3"`
	Valor3      int        `json:"valor3"`
	Liberado3   int        `json:"liberado3"`
	Desde4      int        `json:"desde4"`
	Hasta4      int        `json:"hasta4"`
	Valor4      int        `json:"valor4"`
	Liberado4   int        `json:"liberado4"`
	Desde5      int        `json:"desde5"`
	Hasta5      int        `json:"hasta5"`
	Valor5      int        `json:"valor5"`
	Liberado5   int        `json:"liberado5"`
	Active      int        `json:"active"`
	Reserva     int        `json:"reserva"`
	Author      string     `json:"author"`
	CompanyId   int64      `json:"int64"`
	CreatedDate *time.Time `gorm:"autoCreateTime"`
	UpdatedDate *time.Time `gorm:"autoUpdateTime"`
}

// Resp  response struct
type ProgramResp struct {
	ID          string     `json:"id"`
	Code        string     `json:"code"`
	Name        string     `json:"name"`
	Desde1      int        `json:"desde1"`
	Hasta1      int        `json:"hasta1"`
	Valor1      int        `json:"valor1"`
	Liberado1   int        `json:"liberado1"`
	Desde2      int        `json:"desde2"`
	Hasta2      int        `json:"hasta2"`
	Valor2      int        `json:"valor2"`
	Liberado2   int        `json:"liberado2"`
	Desde3      int        `json:"desde3"`
	Hasta3      int        `json:"hasta3"`
	Valor3      int        `json:"valor3"`
	Liberado3   int        `json:"liberado3"`
	Desde4      int        `json:"desde4"`
	Hasta4      int        `json:"hasta4"`
	Valor4      int        `json:"valor4"`
	Liberado4   int        `json:"liberado4"`
	Desde5      int        `json:"desde5"`
	Hasta5      int        `json:"hasta5"`
	Valor5      int        `json:"valor5"`
	Liberado5   int        `json:"liberado5"`
	Active      int        `json:"active"`
	Reserva     int        `json:"reserva"`
	Author      string     `json:"author"`
	CompanyId   int64      `json:"int64"`
	CreatedDate *time.Time `gorm:"autoCreateTime"`
	UpdatedDate *time.Time `gorm:"autoUpdateTime"`
}

// Create---Req  request struct
type CreateProgramReq struct {
	ID          string     `gorm:"primaryKey;autoIncrement"`
	Code        string     `json:"code"`
	Name        string     `json:"name"`
	Desde1      int        `json:"desde1"`
	Hasta1      int        `json:"hasta1"`
	Valor1      int        `json:"valor1"`
	Liberado1   int        `json:"liberado1"`
	Desde2      int        `json:"desde2"`
	Hasta2      int        `json:"hasta2"`
	Valor2      int        `json:"valor2"`
	Liberado2   int        `json:"liberado2"`
	Desde3      int        `json:"desde3"`
	Hasta3      int        `json:"hasta3"`
	Valor3      int        `json:"valor3"`
	Liberado3   int        `json:"liberado3"`
	Desde4      int        `json:"desde4"`
	Hasta4      int        `json:"hasta4"`
	Valor4      int        `json:"valor4"`
	Liberado4   int        `json:"liberado4"`
	Desde5      int        `json:"desde5"`
	Hasta5      int        `json:"hasta5"`
	Valor5      int        `json:"valor5"`
	Liberado5   int        `json:"liberado5"`
	Active      int        `json:"active"`
	Reserva     int        `json:"reserva"`
	Author      string     `json:"author"`
	CompanyId   int64      `json:"int64"`
	CreatedDate *time.Time `gorm:"autoCreateTime"`
	UpdatedDate *time.Time `gorm:"autoUpdateTime"`
}

type UpdateProgramReq struct {
	ID          string     `json:"-"`
	Code        *string    `json:"code"`
	Name        *string    `json:"name"`
	Desde1      *int       `json:"desde1"`
	Hasta1      *int       `json:"hasta1"`
	Valor1      *int       `json:"valor1"`
	Liberado1   *int       `json:"liberado1"`
	Desde2      *int       `json:"desde2"`
	Hasta2      *int       `json:"hasta2"`
	Valor2      *int       `json:"valor2"`
	Liberado2   *int       `json:"liberado2"`
	Desde3      *int       `json:"desde3"`
	Hasta3      *int       `json:"hasta3"`
	Valor3      *int       `json:"valor3"`
	Liberado3   *int       `json:"liberado3"`
	Desde4      *int       `json:"desde4"`
	Hasta4      *int       `json:"hasta4"`
	Valor4      *int       `json:"valor4"`
	Liberado4   *int       `json:"liberado4"`
	Desde5      *int       `json:"desde5"`
	Hasta5      *int       `json:"hasta5"`
	Valor5      *int       `json:"valor5"`
	Liberado5   *int       `json:"liberado5"`
	Active      *int       `json:"active"`
	Reserva     *int       `json:"reserva"`
	Author      *string    `json:"author"`
	CompanyId   *int64     `json:"int64"`
	CreatedDate *time.Time `gorm:"autoCreateTime"`
	UpdatedDate *time.Time `gorm:"autoUpdateTime"`
}
