package projects

import (
	"fmt"
	"log"
)


func Avg() float64 {
	var x,y float64
    log.Print("enter x")
	fmt.Scanln(&x)

	log.Print("enter y")
	fmt.Scanln(&y)
   if y ==0.0 {
	 panic("number is zero")
   }
	return x/y

}