package dto

type StatusTransaction struct {
	Status string `json:"status"`
}

type PaymentTransaction struct {
	PaymentMethod string `json:"payment_method"`
}
