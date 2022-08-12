package db

type Customer struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Email    string `json:"email"`
	Telephone    string `json:"telephone"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}