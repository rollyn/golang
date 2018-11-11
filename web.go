package main 


import (
		"net/http"
)


func index_handler(w http.ResponseWriter,
 					r *http.Request) {
	http.ServeFile(w, r, "./static/"+r.URL.Path[1:])
	//fmt.Fprintf(w, "<h1>Hello</h1>")
}

func main() {
	
	http.HandleFunc("/", index_handler)

	http.ListenAndServe(":8000", nil)

}