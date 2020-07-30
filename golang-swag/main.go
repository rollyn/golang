// Testing go-swagger generation
//
// The purpose of this application is to test go-swagger in a simple GET request.
//
//     Schemes: http
//     Host: localhost:5000
//     Version: 0.0.1
//     License: MIT http://opensource.org/licenses/MIT
//     Contact: rollyn.moises@gmail.com
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
// swagger:meta
package main

import (
	"encoding/json"
	"net/http"
)

// HTTP status code 200
// swagger:response orderHandler
type orderHandler struct {
	OrderID  string `json:"orderID"`
	Quantity int    `json:"quantity"`
}

// swagger:operation GET /order order order
//
// Returns an order
// ---
// consumes:
// - application/json
// produces:
// - application/json
// responses:
//   '200':
//     description: The Order object
//     type: orderHandler
func (o *orderHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	js, _ := json.Marshal(o)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func main() {

	fs := http.FileServer(http.Dir("./swaggerui"))
	http.Handle("/swaggerui/", http.StripPrefix("/swaggerui/", fs))

	http.Handle("/order", &orderHandler{OrderID: "C001", Quantity: 2})
	http.ListenAndServe(":5000", nil)
}

//creating swagger
//0. go get -u github.com/go-swagger/go-swagger/cmd/swagger
//1. Documentation
//2. swagger generate spec -m -o swagger.json
//3. swagger serve -F=swagger swagger.yaml
