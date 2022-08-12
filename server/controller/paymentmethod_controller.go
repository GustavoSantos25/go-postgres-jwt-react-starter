package controller

import (
	"fmt"
	"net/http"

	"github.com/rogaha/go-postgres-jwt-react-starter/server/db"
	"github.com/rogaha/go-postgres-jwt-react-starter/server/errors"
	"github.com/rogaha/go-postgres-jwt-react-starter/server/utils"

	"github.com/gin-gonic/gin"
)

func CreatePaymentMethod(c *gin.Context){
	var pm db.PaymentMethod
	c.Bind(&pm)
	exists := checkCustomerIdExists(pm.CustomerId)

	valErr := utils.ValidatePaymentMehtod(pm, errors.ValidationErrors)
	if exists != true {
		valErr = append(valErr, "customer doesn't exists")
	}
	fmt.Println(valErr)
	if len(valErr) > 0 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"success": false, "errors": valErr})
		return
	}
	_, err := db.DB.Query(db.CreatePaymentMethodsQuery, pm.Id, pm.MethodType, pm.SuccessfulRegistration, pm.CustomerId)
	errors.HandleErr(c, err)
	c.JSON(http.StatusOK, gin.H{"success": true, "msg": "Payment Method created succesfully"})
}

func GetAllPaymentMethods(c *gin.Context){
	pms := []db.PaymentMethod{}
	rows, _ := db.DB.Query(db.GetAllPaymentMethodsQuery)
	
	for rows.Next(){
		pm := db.PaymentMethod{}

		err := rows.Scan(&pm.Id, &pm.MethodType, &pm.SuccessfulRegistration, &pm.CustomerId, &pm.CreatedAt, &pm.UpdatedAt)
		errors.HandleErr(c, err)
		pms = append(pms, pm)
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "paymentmethods": pms})
}

func checkCustomerIdExists(customerId string) bool {
	rows, err := db.DB.Query(db.CheckCustomerExists, customerId)
	if err != nil {
		return false
	}
	if !rows.Next() {
		return false
	}
	return true
}
