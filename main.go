package main

import (
	"fmt"
	"log"
	"new/projects"
)

func main() {
	// 	projects.projects()
	fmt.Println("hello world")

	//    projects.Slice()
	//    projects.Capacity()
	projects.Map()
	projects.Struct()
	projects.Personal_Input()
	a, _, _, _, _ := projects.Test()
	log.Println(a)

	e,f,g :=projects.Test2()
	log.Println(
		e,f,g,
	)

}

//parameter functions
