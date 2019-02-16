package main

import (
	"log"
	"net/http"
)

var c cow

func main() {
	// c := cow{0}

	http.HandleFunc("/", c.say)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
