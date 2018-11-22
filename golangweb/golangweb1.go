package main 



import ("log"
		"os"
		"fmt"
		"text/template"
)

var tpl *template.Template

func init() {
	fmt.Println("running init....")
	tpl = template.Must(template.ParseFiles("value.gohtml"))
}

func main() {
	


	err := tpl.Execute(os.Stdout, "Rollyn")
	if err != nil{
		log.Fatalln(err)
	}

}