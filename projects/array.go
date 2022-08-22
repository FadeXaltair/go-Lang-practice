package projects

import "fmt"


func Abc(){
  fmt.Println("hello world")

  var matrix [3][3] int
  matrix[0] = [3]int{1,1,1}
  matrix[1] =[3]int{2,2,2}
  matrix[2]=[3]int{3,3,3}

  fmt.Printf( "%v,%T\n", matrix, matrix )


  a :=[3]int{1,2,3}
  b := a
  b[1] =6

  fmt.Println(a)
  fmt.Println(b)

}