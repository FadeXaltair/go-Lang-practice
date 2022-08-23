package projects

import (
	"fmt"
	"log"
)


func Switch() {
	var day int 
	log.Println("enter your day : ", )
	fmt.Scan(&day)

	switch day {
	case 1:
		log.Println("sunday")
	
   case 2: 
	log.Println("mon")


    case 3:
	log.Println("tue")


    case 4:
	log.Println("wed")


    case 5:
	log.Println("thrs")


    case 6:
	log.Println("fri")

     case 7:
	log.Println("sat")

	default :
	log.Println("no found")
	
    }
}