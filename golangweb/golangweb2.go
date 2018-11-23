package main 



import ("log"
		"os"
		"fmt"
		"text/template"
)

var tpl *template.Template


type person struct {
	lname string
	fname string
}

func init() {
	fmt.Println("running init....")
	tpl = template.Must(template.ParseGlob("type.gohtml"))
}

func main() {
	
	fruits := []string{"banana","apple","guava"}

	err := tpl.Execute(os.Stdout, fruits)
	if err != nil{
		log.Fatalln(err)
	}

}