package main

import "fmt"

func add(x, y float64) float64 {
	return x + y
}

func multiple(a, b string) (string, string) {
	return b, a
}

func main() {

	//var num1, num2 float64 = 5.6,9.5
	num1, num2 := 5.6, 9.5

	fmt.Println( add(num1, num2))
	
	a, b := multiple("hello","world")
	fmt.Println(a,b)
}