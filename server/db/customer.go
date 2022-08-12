package db

type Customer struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Email    string `json:"email"`
	Telephone    string `json:"telephone"`
}