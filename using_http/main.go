package main 


import(
"html/template"
"net/http"
"log")

type hotdog int

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func (h hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}
	w.Header().Set("Content-Type","text/plain; charset=utf-8")
	tpl.ExecuteTemplate(w, "index.gohtml", r.PostForm)
}

func main() {
	
	var d hotdog
	http.ListenAndServe(":7070", d)
}