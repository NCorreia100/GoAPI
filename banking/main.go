package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"log"
	"net/http"
)

type Customer struct{
	Name string `json:"name" xml:"name"`
	Age int16 `json:"age" xml:"age"`
	IsFemale bool `json:"isFemale" xml:"isFemale"`
}

func greet(w http.ResponseWriter,  r *http.Request){
	fmt.Fprint(w, "hello Anh!!")
}

func getAllCustomers(w http.ResponseWriter,  r *http.Request){
	customers:=[] Customer{
		{Name: "Anh", IsFemale: true, Age: 29},
		{Name: "Ram", IsFemale: false, Age: 24},
	}
	if(r.Header.Get("Content-Type")=="application/xml"){
		w.Header().Add("Content-Type","application/xml")
		xml.NewEncoder(w).Encode(customers)
	}else{
		w.Header().Add("Content-Type","application/json")
		json.NewEncoder(w).Encode(customers)
	}
}

func main() {
	http.HandleFunc("/greet",greet)
	http.HandleFunc("/customers",getAllCustomers)

	log.Fatal(http.ListenAndServe("localhost:8000",nil))
}