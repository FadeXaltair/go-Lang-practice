package projects

import (
	"fmt"
	"net/mail"
)

func ValidMailAddress() bool{
	var email string
	fmt.Print("Enter your email : ")
	fmt.Scan(&email)
   _,err := mail.ParseAddress(email)
   return err == nil

}