package main 



import ("log"
		"os"
		"fmt"
		"text/template"
)

var tpl *template.Template


type person struct {
	Lname string
	Fname string
}

func (p person) Speak() string {
	// fmt.Println(p.Fname, " speaks....")
	return p.Fname + " speaks...."
}

func init() {
	fmt.Println("running init....")
	tpl = template.Must(template.ParseGlob("type2.gohtml"))
}

func main() {
	
	p := person{"moises","rollyn"}

	err := tpl.Execute(os.Stdout, p)
	if err != nil{
		log.Fatalln(err)
	}

}