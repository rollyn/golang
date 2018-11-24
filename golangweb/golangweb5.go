package main 



import ("log"
		"os"
		"text/template"
		"strings"
		"time"
)

var tpl *template.Template


var fm = template.FuncMap{
	"uc": strings.ToUpper,
	"ft": firstThree,
	"fDateYr": monthDayYear,
}

func firstThree(s string) string {
	s = strings.TrimSpace(s)
	s = s[:3]
	return s
}

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("format.gohtml"))
}

func monthDayYear(t time.Time) string {
	return t.Format("01-02-2006")
}

type person struct {
	Lname string
	Fname string
}

func main() {

	// time.Now()
	err := tpl.ExecuteTemplate(os.Stdout, "format.gohtml", "acidrain")
	if err != nil{
		log.Fatalln(err)
	}

}