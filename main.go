package main

import (
	"log"
	"net/http"
)

var c cow

func init() {
	log.Printf("%s version %s initialized", APPNAME, VERSION)
}

func main() {

	http.HandleFunc("/", logging(c.say))
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func logging(h http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		c.requests++
		ua := r.UserAgent()

		log.Printf("%v requested: %v host: %v, user-agent: %s", c.requests, r.RequestURI, r.RemoteAddr, ua)
		h(w, r)
	}

}
