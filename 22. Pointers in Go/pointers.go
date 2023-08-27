package main

import "fmt"

func main() {
	a := "Aadarsh"
	fmt.Println(&a)

	b := &a
	fmt.Printf("b is of type %T with value %v\n", b, b)
	fmt.Printf("address of b: %p\n", &b)

	var ptr1 *float64
	_ = ptr1 //to handle the unused variable

	//s := 50.34
	//ptrS := &s
	//ptr1 = ptrS

	p := new(int) //will allocate new memory
	x := 100
	p = &x
	fmt.Printf("p is of type %T with value %v\n", p, p)

	fmt.Println("value of x: ", *p)
	*p = 49
	fmt.Println("value of x: ", *p)
}
