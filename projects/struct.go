package projects

import "fmt"

type variable struct {
	number     int
	name       string
	companions []string
}


func Struct() {

	a := variable{
		number: 1,
		name:   "hitesh",
		companions: []string{
			"hello", "hitesh",
		},
	}

	fmt.Println(a.name)

}

func Test() (string, string, string, string, string) {
	return "hello", "hello2", "hello3", "hello4", "hello4"
}

func Test2() (string , int, float32){
	return "hello" ,15 , 25.66
}
