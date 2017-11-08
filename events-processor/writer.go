package processor

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"os"
	"strconv"
	"time"
)

func WriteProcessedEvents() {
	go func() {
		for range time.Tick(5000 * time.Millisecond) {
			fmt.Println("ping")
			write()
		}
	}()
}

func write() {
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

	fileName := "conversions" + strconv.Itoa(int(time.Now().UnixNano()))
	tmpPath := "./tmp/" + fileName
	path := "./conversions/" + fileName

	f, err := os.Create(tmpPath)
	check(err)
	defer f.Close()

	var buffer bytes.Buffer

	for key, value := range double {
		buffer.WriteString(key + ":" + strconv.Itoa(value) + "\n")
	}

	_, err = f.WriteString(buffer.String())
	check(err)

	err = os.Rename(tmpPath, path)
	check(err)

	return
}
