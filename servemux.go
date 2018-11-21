package main 


import("net/http"
)

func Index(w http.ResponseWriter,
 					r *http.Request) {
	http.ServeFile(w, r, "./static/"+r.URL.Path[1:])
}

func main() {
	mux := http.NewServeMux()

	//th := http.HandlerFunc(Index)
	//mux.Handle("/", th)

	mux.HandleFunc("/", Index)
	
	http.ListenAndServe(":3000", mux)

}