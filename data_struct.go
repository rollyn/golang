package main 


import "fmt"


type person struct {
	fname string
	lname string
}

type secretAgent struct {
	person
	licenseToKill bool
}

func (p person) speak() {
	fmt.Println(p.lname, " speak")
}

func (sa secretAgent) speak() {
	fmt.Println(sa.fname, "Run...")
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func main() {
	
//	xi := []int{1,2,3,4}
	// fmt.Println(xi)

	// m := map[string]int{
	// 	"Ice":20,
	// 	"Eth":10,
	// }
	// fmt.Println(m["Ice"])

	p1 := person{fname: "Rollyn", lname: "Moises"}
	// fmt.Println(p1, p1.fname)

	// p1.speak()


	sa := secretAgent{
		person{"Rollyn","Moises"},true,
	}
	// fmt.Println(sa)

	// sa.speak()
	// sa.person.speak()

	saySomething(p1)
	saySomething(sa)


}

