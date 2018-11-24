package main 



import ("log"
		"os"
		"fmt"
		"text/template"
)

var tpl *template.Template

func init() {
	fmt.Println("running init....")
	tpl = template.Must(template.ParseGlob("*.gohtml"))
}

func main() {
	


	err := tpl.ExecuteTemplate(os.Stdout, "index.gohtml", 42)
	if err != nil{
		log.Fatalln(err)
	}

}