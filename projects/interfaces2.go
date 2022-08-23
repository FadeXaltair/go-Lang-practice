package projects

import "log"

type SalaryCalculator interface{
	calculateSalary() int
}


type Permanent struct{
	basicpay int
	pf int
}

type Contract struct{
	basicpay int
}


func(p Permanent) calculateSalary() int{
	return p.basicpay + p.pf
}

func( c Contract) calculateSalary() int {
	return c.basicpay
}


func Total_expense(s [] SalaryCalculator){

	expenses :=0

	for _,v := range s {
		expenses += v.calculateSalary()
	}
  log.Println(" total expenses : ", expenses)
}


func Interfacer() {  
    p1 := Permanent{
        basicpay: 5000,
        pf:20,
    }
    p2 := Permanent{
        basicpay: 6000,
        pf:30,
    }
    c1 := Contract{
        basicpay: 3000,
    }
    employees := []SalaryCalculator{p1, p2, c1}
    Total_expense(employees)

}