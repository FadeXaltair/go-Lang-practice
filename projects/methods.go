package projects

import (
	"log"
	"math"
)


type rectangle struct{
	length int
	breadth int
}

type circle struct {
	radius float64
}


func (r rectangle) Area() int {
	return r.length * r.breadth
}

func (c circle) Area() float64 {
	return math.Pi *c.radius*c.radius
}



func Cir_Rec() {
   a := rectangle{
	length : 3,
	breadth: 4,
   }

   b := circle{
	radius: 3.00,
   }

  log.Println("area of rectangele : ", a.Area())

  log.Println("AREA of circle : " , b.Area())
}