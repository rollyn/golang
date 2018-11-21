package main 


import("fmt"
		"encoding/json"
		"net/http"
)


type Person struct {
	Name string `json:"name"`
	Age int `json:"age"`
}

func Index(w http.ResponseWriter, r *http.Request) {

	p := Person{Name: "Rollyn", Age:20}
	data, err := json.Marshal(p)

	if err != nil {
		panic("Opps")
	}
	fmt.Fprint(w, string(data))
} 
func main() {
	
	

	http.HandleFunc("/", Index)
	http.ListenAndServe(":8000", nil)

}