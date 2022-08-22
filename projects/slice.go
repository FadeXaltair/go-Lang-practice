package projects

import "fmt"

func Slice(){

	a :=[]int{1,2,3,4,5,6,7,8,9}
	b :=a[:]
	c :=a[3:]
	d :=a[:5]
	e :=a[3:6]

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)


}

func Capacity(){

	a := []int{}
	fmt.Println(a)
	fmt.Printf("length size, %v\n", len(a))
	fmt.Printf("capacaity size, %v\n", cap(a))

	a = append(a,1)
	fmt.Println(a)
	fmt.Printf("length size: %v\n", len(a))
	fmt.Printf("capacaity size: %v\n", cap(a))


	a = append(a,2,3,4,5)
	fmt.Println(a)
	fmt.Printf("length size: %v\n", len(a))
	fmt.Printf("capacaity size: %v\n", cap(a))

	b :=[]int{1,2,3,4,5}
	b = append(b[:2],b[3:]...)
	fmt.Println(b)
  

}