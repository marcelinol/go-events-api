package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/marcelinol/go-events-api/events-processor"
)

func main() {

	http.HandleFunc("/event", processor.EnqueueEvent)

	fmt.Println("Listening on port :8080")
	processor.ProcessEvents()
	processor.WriteProcessedEvents()

	log.Fatal(http.ListenAndServe(":8080", nil))
}
