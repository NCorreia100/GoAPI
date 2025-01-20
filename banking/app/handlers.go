package app

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"

	"github.com/NCorreia100/GoAPI/banking/service"
	"github.com/gorilla/mux"
)

type Customer struct {
	Name     string `json:"name" xml:"name"`
	Age      int16  `json:"age" xml:"age"`
	IsFemale bool   `json:"isFemale" xml:"isFemale"`
}

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello Anh!!")
}

type CustomerHandlers struct {
	service service.CustomerService
}

func (ch *CustomerHandlers) getAllCustomers(w http.ResponseWriter, r *http.Request) {
	// customers := []Customer{
	// 	{Name: "Anh", IsFemale: true, Age: 29},
	// 	{Name: "Ram", IsFemale: false, Age: 24},
	// }

	customers, _ := ch.service.GetAllCustomers()

	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(customers)
	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customers)
	}
}

func getCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprint(w, vars["id"])
}
