package main 


import ("net/http"
   "log"
	L "./lib"	
)


func main() {

	http.HandleFunc("/", L.Index())
	log.Println("Listening...")
	http.ListenAndServe(":8000",nil)	
}