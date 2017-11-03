package main

import (
	"encoding/json"
	"fmt"
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

type Processor struct {
	Channel chan string
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
	processor := Processor{make(chan string)}
	StartProcessing(processor.Channel)
	processor.Channel <- contact.Email
}

func StartProcessing(channel chan string) {
	go func() {
		for {
			select {
			case email := <-channel:
				fmt.Println("Received conversion")
				fmt.Println(email)
				WriteToFile(email)
			}
		}
	}()
}

func WriteToFile(email string) {
	file, err := os.Create("./../go-events-processor/tmp/conversions" + strconv.Itoa(int(time.Now().UnixNano())))
	check(err)

	_, err = file.WriteString(email)
	check(err)

	defer file.Close()
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/event", ProcessEvent).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", router))
}
