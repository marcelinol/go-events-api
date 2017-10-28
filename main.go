package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Contact struct {
	Email string `json:"email"`
}

func CreateContactEndpoint(w http.ResponseWriter, req *http.Request) {
	var contact Contact
	_ = json.NewDecoder(req.Body).Decode(&contact)
	fmt.Println(contact)
	json.NewEncoder(w).Encode(contact)
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/event", CreateContactEndpoint).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", router))
}
