package projects

import (

	"fmt"

)


func Personal_Input() {
	var x,y int
	total := 0
    fmt.Printf("enter first number : ")
	fmt.Scanln(&x)

	fmt.Printf("enter second number : " )
	fmt.Scanln(&y)

	total = x+y

	fmt.Println("your sum is :", total)


}