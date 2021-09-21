package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	address := flag.String("address", ":8081", "address")

	flag.Parse()

	server := &http.Server{Addr: *address, Handler: handler()}

	log.Println("starting server", *address)

	err := server.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}
}

func handler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write(readFile())
	})
}

func readFile() []byte {
	// Open our jsonFile
	jsonFile, err := os.Open("response.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	//var result map[string]interface{}
	//json.Unmarshal([]byte(byteValue), &result)

	return []byte(byteValue)
}
