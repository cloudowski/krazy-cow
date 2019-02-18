package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

type cow struct {
	requests int
	name     string
}

// v----------
const asciicow string = `
           (    )
            (oo)
   )\.-----/(O O)
  # ;       / u
    (  .   |} )
     |/ ".;|/;
     "     " "
`

func (c *cow) init() {

	if name, err := os.Hostname(); err != nil {
		log.Fatalln("Failed to get cow name (read hostname)")
	} else {
		c.name = name
	}

}

func (c *cow) say(w http.ResponseWriter, r *http.Request) {

	msg := fmt.Sprintf("\"%s\"", cowconf.Get("cow.say"))
	fmt.Fprintf(w, "%15s %s %s", " ", msg, asciicow)

}

func (c *cow) setfree(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Moooooooo! (cow %s has been set free)", c.name)
	go func() {
		time.Sleep(time.Second * 3)
		log.Fatalln("Cow has been set free!")
	}()

}
