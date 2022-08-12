package db

//CreateCustomersTable //creates customers table
func CreateCustomersTable() {
	DB.Query(`
		CREATE TABLE IF NOT EXISTS customers( id VARCHAR(100) PRIMARY KEY, name VARCHAR (100) NOT NULL, email VARCHAR (355) UNIQUE NOT NULL, telephone VARCHAR(100) NOT NULL,created_on TIMESTAMP NOT NULL default current_timestamp,updated_at TIMESTAMP NOT NULL default current_timestamp )`,
	)
}