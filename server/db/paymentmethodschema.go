package db

//CreatePaymentMethodsTable //creates payment methods table
func CreatePaymentMethodsTable() {
	DB.Query(`
		CREATE TABLE IF NOT EXISTS paymentmethods( id VARCHAR(100) PRIMARY KEY, method_type VARCHAR(100) NOT NULL, successful_registration BOOLEAN NOT NULL, customer_id VARCHAR(100) NOT NULL, created_on TIMESTAMP NOT NULL default current_timestamp,updated_at TIMESTAMP NOT NULL default current_timestamp, CONSTRAINT chk_method_type CHECK (method_type IN ('card', 'debit', 'paypal')), CONSTRAINT fk_customer FOREIGN KEY(customer_id) REFERENCES customers(id) )`,
	)
}
