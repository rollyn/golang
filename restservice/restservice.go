package main 


import("fmt"
		"net/http"
		"github.com/julienschmidt/httprouter"
		"./controllers"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	fmt.Fprint(w, "Welcome\n")
}


func main() {
	
	r := httprouter.New()

	uc := controllers.NewUserController()

	// Get a user resource
	r.GET("/user/:id", uc.GetUser)

	http.ListenAndServe("localhost:3000",r)
}