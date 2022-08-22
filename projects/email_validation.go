package projects

import (
	"fmt"
	"log"
	"net/mail"
)

func ValidMailAddress() {
	var email string
	fmt.Print("Enter your email : ")
	fmt.Scan(&email)
   _,err := mail.ParseAddress(email)
   if err != nil{
	log.Print("invalid email address :", email)
   }else{
   log.Print("valid email " , email)
}

}