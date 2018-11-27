package main 


import("fmt")


type Information interface {
	General()
	Inventory()
	Attributes()
}

type Product struct {
	Name string
	Description string
}

func (p Product) General() {
	fmt.Println(p.Name)
}

func (p Product) Inventory() {
	fmt.Println(p.Description)
}




func main() {

	p := Product{Name:"Gold",Description:"24 Karat"}

	p.General()
	p.Inventory()
	
}