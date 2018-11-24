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

type doubleZero struct {
	person
	LicenseToKill bool
}

func init() {
	fmt.Println("running init....")
	tpl = template.Must(template.ParseFiles("zero.gohtml"))
}

func main() {
	

	p := doubleZero{ person{"bond", "james"}, true}
	err := tpl.Execute(os.Stdout, p)
	if err != nil{
		log.Fatalln(err)
	}

}