package timezone

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func StartServer() {
	router:= mux.NewRouter()
	router.HandleFunc("/api/time", getTime).Methods(http.MethodGet)
	

	log.Fatal(http.ListenAndServe("localhost:8000", router))
}
