package main 


import("fmt")

type Person struct {
	Fname string
	Lname string
}

func main() {

	// using a pointer reference allow us to 
	// copy struct that is inexpensive
	
	p := &Person{Fname:"ice",Lname:"man"}

	fmt.Println(p.Fname)

}