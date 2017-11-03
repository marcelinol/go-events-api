package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type Contact struct {
	Email string `json:"email"`
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func ProcessEvent(w http.ResponseWriter, req *http.Request) {
	var contact Contact
	err := json.NewDecoder(req.Body).Decode(&contact)
	check(err)
	WriteToFile(contact)
}

func WriteToFile(contact Contact) {
	file, err := os.Create("./../go-events-processor/tmp/conversions" + strconv.Itoa(int(time.Now().UnixNano())))
	check(err)

	_, err = file.WriteString(contact.Email)
	check(err)

	defer file.Close()
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/event", ProcessEvent).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", router))
}
