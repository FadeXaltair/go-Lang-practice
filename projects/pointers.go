package projects

import (
	"log"
)


func Pointers(){
	i,j := 40,70

	p :=&i
	log.Println(*p)

	*p = 21

	log.Println(i)

	p = &j
	*p = *p/35

	log.Println(j)

}