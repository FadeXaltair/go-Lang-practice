package projects

import (
	"fmt"
)

func Loop(){
    // i :=0
	// for i,j :=0,0 ; i< 6; i,j= i+1,j+1 {
	// 	fmt.Println(i,j)
	// }


	// for i<5 {
	// 	i ++
    //   fmt.Print( i)
	// }

		//infinte ;loop syntax
	// for{
	// 	fmt.Println(i)
	// 	i++
    //    if (i ==6){
	// 	break
	//    }
	// }

	s :=[]int{ 2,3,4}

	for k,v := range s {
		fmt.Println(k,v)
	}

	a := "Hello world"
	for k,v := range a {
		fmt.Println(k, string(v))
	}
}