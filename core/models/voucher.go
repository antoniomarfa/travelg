package models

type Voucher struct {
	ID        string `json:"_id,omitempty"`
	SaleId    int64  `json:"sale_id"`
	Voucher   string `json:"voucher"`
	Used      int    `json:"used"`
	CompanyId int64  `json:"company_id"`
}

// Resp  response struct
type VoucherResp struct {
	ID        string `json:"id"`
	SaleId    int64  `json:"sale_id"`
	Voucher   string `json:"voucher"`
	Used      int    `json:"used"`
	CompanyId int64  `json:"company_id"`
}

// Create---Req  request struct
type CreateVoucherReq struct {
	ID        string `gorm:"primaryKey;autoIncrement"`
	SaleId    int64  `json:"sale_id"`
	Voucher   string `json:"voucher"`
	Used      int    `json:"used"`
	CompanyId int64  `json:"company_id"`
}

type UpdateVoucherReq struct {
	ID        string `json:"-"`
	SaleId    int64  `json:"sale_id"`
	Voucher   string `json:"voucher"`
	Used      int    `json:"used"`
	CompanyId int64  `json:"company_id"`
}
