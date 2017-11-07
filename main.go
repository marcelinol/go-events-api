package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/marcelinol/go-events-api/events-processor"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/event", processor.EnqueueEvent).Methods("POST")
	fmt.Println("Listening on :8080")
	processor.ProcessEvents()
	processor.WriteProcessedEvents()

	log.Fatal(http.ListenAndServe(":8080", router))
}
