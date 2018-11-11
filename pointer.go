package main 

import ("fmt"

)

func main() {
	
	x := 5
	a := &x

	fmt.Println(a)
	fmt.Println(*a)

	*a = 10
	fmt.Println(x)
}
