package projects

import "log"

type Employee interface{
	PrintName(name string)
	PrintSalary(basic,tax int) int
}


type Emp int 


func (e Emp) PrintName(name string){
	log.Print("employee ud : ",e)
	log.Print("employee name : ",name)
}

func (e Emp) PrintSalary(basic,tax int) int {
	var sal = (basic * tax)/100
	return sal-basic
}

func Caller() {
	e1 := Emp(1)
	e1.PrintName("hitesh")
	log.Println(e1.PrintSalary(5000,1000))
}