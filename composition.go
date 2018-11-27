package main 

import("fmt")


type Person struct {
	Name string
}

type Player struct {
	Person
	Number string
}
func main() {
	
	p := Player{}
	p.Name = "acid"
	p.Number = "27"

	fmt.Println(p)

}