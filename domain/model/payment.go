package model

type Payment struct {
	ID            int64  `gorm:"primary_key;not_null;auto_increment"`
	PaymentName   string `json:"payment_name"`
	PaymentSID    string `json:"payment_s_id"`
	PaymentStatus bool   `json:"payment_status"`
	PaymentImage  string `json:"payment_image"`
}
