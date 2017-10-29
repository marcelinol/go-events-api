package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

type Contact struct {
	Email string `json:"email"`
}

func CreateContactEndpoint(w http.ResponseWriter, req *http.Request) {
	var contact Contact
	err := json.NewDecoder(req.Body).Decode(&contact)

	check(err)

	fmt.Println(contact)
	fmt.Println(contact.Email)

	f, err := os.Create("./tmp/file")
	check(err)
	defer f.Close()

	_, err = f.WriteString(contact.Email)
	check(err)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/event", CreateContactEndpoint).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", router))
}
