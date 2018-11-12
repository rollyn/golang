package main 


import "fmt"

func main() {
	
	// for i:=0;i<10;i++{
	// 	fmt.Println(i)
	// }

	i := 0
	for i<10 {
		fmt.Println(i)
		i++
	}

	x := 5
	for {
		fmt.Println("Running ",x)
		x += 3
		if x > 20 {
			break
		}
	}
}