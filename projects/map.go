package projects

import "fmt"

func Map(){
   //if you dont have the value we can just initialise them

	// a := make(map[string]int)
	a := map[string]int{
		"value 1" : 1,
		"value 2" : 2,
	}

	fmt.Println(a["value 1"])

	//deletion
	delete(a, "value 1")
	//addition
	a["value 3"] = 3

	verify, ok := a["value 2"]


	fmt.Println(verify, ok)
}