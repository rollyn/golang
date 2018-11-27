package main 



import("fmt")

type Person struct {
	name string
}

func (p *Person) Speak() string {
	return p.name + " Speak..."
}

func main() {
	
	p := &Person{"acid"}

	fmt.Println( p.Speak() )
}