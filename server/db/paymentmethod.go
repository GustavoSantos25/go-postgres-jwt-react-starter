package db

type PaymentMethod struct {
	Id string `json:"id"`
	MethodType string `json:"method_type"`
	SuccessfulRegistration bool `json:"successful_registration"`
	CustomerId string `json:"customer_id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}