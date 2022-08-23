package projects

import (
	"fmt"
	"log"
)


func Sum(values ...int){
	log.Println(values)

	result :=0

	for _,v :=range values{
		result += v
	}

	log.Println("sum is : " , result)
}

func Sum2(values ...int) (result int) {

	log.Println(values)

	for _,v := range values{
		result += v
	}
  return
}

func Divide(a,b float64) (float64, error){

	if b == 0.0{
		return 0.0, fmt.Errorf("cannot divide zero")
	}
	return a/b , nil
}

//