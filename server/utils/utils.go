package utils

import (
	"regexp"

	"github.com/rogaha/go-postgres-jwt-react-starter/server/db"
)

const (
	emailRegex = `^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`
	telephoneRegex = `\(?([0-9]{3})\)?([ .-]?)([0-9]{3})\2([0-9]{4})`
)

// ValidateUser returns a slice of string of validation errors
func ValidateUser(user db.Register, err []string) []string {
	emailCheck := regexp.MustCompile(emailRegex).MatchString(user.Email)
	if emailCheck != true {
		err = append(err, "Invalid email")
	}
	if len(user.Password) < 4 {
		err = append(err, "Invalid password, Password should be more than 4 characters")
	}
	if len(user.Name) < 1 {
		err = append(err, "Invalid name, please enter a name")
	}

	return err
}

func ValidateCustomer(customer db.Customer, err []string) []string {
	emailCheck := regexp.MustCompile(emailRegex).MatchString(customer.Email)
	// telephoneCheck := regexp.MustCompile(telephoneRegex).MatchString(customer.Telephone)
	if emailCheck != true {
		err = append(err, "Invalid email")
	}
	// if telephoneCheck != true {
	// 	err = append(err, "Invalid telephone")
	// }
	if len(customer.Name) < 1 {
		err = append(err, "Invalid name, please enter a name")
	}

	return err
}

func ValidatePaymentMehtod(pm db.PaymentMethod, err []string) []string {
	if !containsString([]string{"card", "debit", "paypal"}, pm.MethodType) {
		err = append(err, "Invalid method type")
	}
	return err
}

func containsString(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}
