package main

import (
	"bytes"
	"encoding/gob"
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
	Channel    chan string
	Aggregator map[string]int
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
	processor.Channel <- contact.Email
}

func StartProcessing(channel chan string) {
	go func() {
		for {
			select {
			case email := <-channel:
				fmt.Println("Conversion Received")
				EmailMapper(email)
			}
		}
	}()
}

func StartWriting(channel chan string) {
	go func() {
		for range time.Tick(5000 * time.Millisecond) {
			fmt.Println("ping")
			Write()
		}
	}()
}

func Write() {
	if len(processor.Aggregator) < 1 {
		return
	}
	double := processor.Aggregator
	processor.Aggregator = make(map[string]int)
	fmt.Println("")
	b := new(bytes.Buffer)

	e := gob.NewEncoder(b)

	// Encoding the map
	err := e.Encode(double)
	if err != nil {
		panic(err)
	}

	var decodedMap map[string]int
	d := gob.NewDecoder(b)

	// Decoding the serialized data
	err = d.Decode(&decodedMap)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%#v\n", decodedMap)

	f, err := os.Create("./../go-events-processor/tmp/conversions" + strconv.Itoa(int(time.Now().UnixNano())))
	check(err)
	defer f.Close()

	var buffer bytes.Buffer

	for key, value := range double {
		buffer.WriteString(key + ":" + strconv.Itoa(value) + "\n")
	}

	_, err = f.WriteString(buffer.String())
	check(err)

	return
}

func EmailMapper(email string) {
	_, emailExist := processor.Aggregator[email]
	if emailExist {
		processor.Aggregator[email]++
		fmt.Printf("email %s converted with count %d\n", email, processor.Aggregator[email])
	} else {
		processor.Aggregator[email] = 1
		fmt.Printf("email %s converted with count %d\n", email, 1)
	}
}

func WriteToFile(email string) {
	file, err := os.Create("./../go-events-processor/tmp/conversions" + strconv.Itoa(int(time.Now().UnixNano())))
	check(err)

	_, err = file.WriteString(email)
	check(err)

	defer file.Close()
}

var processor = Processor{make(chan string), make(map[string]int)}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/event", ProcessEvent).Methods("POST")
	fmt.Println("Listening on :8080")
	StartProcessing(processor.Channel)
	StartWriting(processor.Channel)

	log.Fatal(http.ListenAndServe(":8080", router))
}
