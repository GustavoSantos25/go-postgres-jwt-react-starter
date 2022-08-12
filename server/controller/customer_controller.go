package controller

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/rogaha/go-postgres-jwt-react-starter/server/db"
	"github.com/rogaha/go-postgres-jwt-react-starter/server/errors"
	"github.com/rogaha/go-postgres-jwt-react-starter/server/utils"

	"github.com/gin-gonic/gin"
)

func CreateCustomer(c *gin.Context){
	var customer db.Customer
	c.Bind(&customer)
	exists := checkCustomerExists(customer)

	valErr := utils.ValidateCustomer(customer, errors.ValidationErrors)
	if exists == true {
		valErr = append(valErr, "email already exists")
	}
	fmt.Println(valErr)
	if len(valErr) > 0 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"success": false, "errors": valErr})
		return
	}
	_, err := db.DB.Query(db.CreateCustomerQuery, customer.Id, customer.Name, customer.Email, customer.Telephone)
	errors.HandleErr(c, err)
	c.JSON(http.StatusOK, gin.H{"success": true, "msg": "Customer created succesfully"})
}

func GetAllCustomers(c *gin.Context){
	customers := []db.Customer{}
	rows, _ := db.DB.Query(db.GetAllCustomersQuery)
	
	for rows.Next(){
		customer := db.Customer{}

		err := rows.Scan(&customer.Id, &customer.Name, &customer.Email, &customer.Telephone, &customer.CreatedAt, &customer.UpdatedAt)
		errors.HandleErr(c, err)
		customers = append(customers, customer)
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "customers": customers})
}

func GetCustomerById(c *gin.Context){
	customer := db.Customer{}
	row := db.DB.QueryRow(db.GetCustomerByIdQuery, c.Param("id"))

	err := row.Scan(&customer.Id, &customer.Name, &customer.Email, &customer.Telephone, &customer.CreatedAt, &customer.UpdatedAt)
	if err == sql.ErrNoRows {
		fmt.Println(sql.ErrNoRows, "err")
		c.JSON(http.StatusNotFound, gin.H{"success": false, "msg": "invalid customer id"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "customers": customer})
}


func checkCustomerExists(customer db.Customer) bool {
	rows, err := db.DB.Query(db.CheckCustomerExists, customer.Id)
	if err != nil {
		return false
	}
	if !rows.Next() {
		return false
	}
	return true
}

