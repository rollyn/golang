package main 

import "fmt"

func main() {
	grades := make(map[string]float32)

	grades["john"] = 42
	grades["peter"] = 55
	grades["luke"] = 45

	fmt.Println(grades)

	JohnGrade := grades["john"]
	fmt.Println(JohnGrade)

	delete(grades,"john")
	
	fmt.Println(grades)
	grades["john"] = 55
	grades["john"] = 65
	for k,v := range grades {
		fmt.Println(k,v)
	}

	// remove var if using :=
	ages := make(map[string]float32)  // need to do make since its nil

	ages["jude"] = 10

	fmt.Println(ages)

}