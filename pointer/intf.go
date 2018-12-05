package main 


import("fmt")


type Person struct {
	Fname string
	Lname string
}

func (p *Person) Name() string {
	return fmt.Sprintf("%s %s", p.Fname, p.Lname)
}

type Namer interface {
	Name() string
}

func Greet(n Namer) string {
	return  fmt.Sprintf("Dear %s", n.Name())
}

func main() {
	p := &Person{"ice","man"}
	fmt.Println( Greet(p) )
}