package main 


import(
"github.com/julienschmidt/httprouter"
"html/template"
"net/http"
"log"
)


var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("views/*"))
}


func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	err := tpl.ExecuteTemplate(w, "index.gohtml",nil)
	HandleError(w, err)
}

func about(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	err := tpl.ExecuteTemplate(w, "about.gohtml",nil)
	HandleError(w, err)
}


func HandleError(w http.ResponseWriter, err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {

	mux := httprouter.New()
	mux.GET("/", index)
	mux.GET("/about", about)	
	http.ListenAndServe(":7070", mux)
}


