package main 



import ("log"
		"os"
		"text/template"
		"strings"
)

var tpl *template.Template


var fm = template.FuncMap{
	"uc": strings.ToUpper,
	"ft": firstThree,
}

func firstThree(s string) string {
	s = strings.TrimSpace(s)
	s = s[:3]
	return s
}

func init() {
	tpl = template.Must(template.New("funcmap.gohtml").Funcs(fm).ParseFiles("funcmap.gohtml"))
}

type person struct {
	Lname string
	Fname string
}

func main() {

	p := person{"iceage ","manuever"}

	err := tpl.Execute(os.Stdout, p)
	if err != nil{
		log.Fatalln(err)
	}

}