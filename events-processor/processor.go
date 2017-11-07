package processor

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Contact struct {
	Email string `json:"email"`
}

type Processor struct {
	Channel    chan string
	Aggregator map[string]int
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

var processor = Processor{make(chan string), make(map[string]int)}

func EnqueueEvent(w http.ResponseWriter, req *http.Request) {
	var contact Contact
	err := json.NewDecoder(req.Body).Decode(&contact)
	check(err)
	processor.Channel <- contact.Email
}

func ProcessEvents() {
	go func() {
		for {
			select {
			case email := <-processor.Channel:
				fmt.Println("Conversion Received")
				emailMapper(email)
			}
		}
	}()
}

func emailMapper(email string) {
	_, emailExist := processor.Aggregator[email]
	if emailExist {
		processor.Aggregator[email]++
		fmt.Printf("email %s converted with count %d\n", email, processor.Aggregator[email])
	} else {
		processor.Aggregator[email] = 1
		fmt.Printf("email %s converted with count %d\n", email, 1)
	}
}
