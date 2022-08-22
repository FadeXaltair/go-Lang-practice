package projects

import "fmt"

// handleing panic with recover
func HandlePanaic(){
	a := recover()
	if a != nil {
		fmt.Println("recover", a)
	}
}

// adding defer and panic 
func Testing(num1, num2 int){
	defer HandlePanaic()
  if num2 == 0{
	panic("end the function as the value is zero")
  } else{
	result := num1/num2
	fmt.Println("result is : ", result)
  }

}