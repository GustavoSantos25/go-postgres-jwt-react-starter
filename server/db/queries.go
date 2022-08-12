package db

const (
	CheckUserExists = `SELECT id from users WHERE email = $1`
	LoginQuery      = `SELECT * from users WHERE email = $1`
	CreateUserQuery = `INSERT INTO users(id,name,password,email) VALUES (DEFAULT, $1 , $2, $3);`
	CheckCustomerExists = `SELECT * FROM customers WHERE id = $1`
	CreateCustomerQuery = `INSERT INTO customers(id,name,email,telephone) VALUES ($1, $2, $3, $4);`
	GetAllCustomersQuery = `SELECT * FROM customers`
	GetCustomerByIdQuery = `SELECT * FROM customers WHERE id = $1`
	CreatePaymentMethodsQuery = `INSERT INTO paymentmethods(id, method_type, successful_registration, customer_id) VALUES ($1, $2 , $3, $4)`
	GetAllPaymentMethodsQuery = `SELECT * FROM paymentmethods`
	GetPaymentMethodByIdQuery = `SELECT * FROM paymentmethods WHERE id = $1`
	GetCustomerPaymentMethodsQuery = `SELECT * FROM paymentmethods WHERE customer_id = $1`
	GetCustomerPaymentMethodsCountQuery = `SELECT COUNT(*) FROM paymentmethods WHERE customer_id = $1`
)
