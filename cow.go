package main

import (
	"fmt"
	"log"
	"net/http"
)

type cow struct {
	replies int
}

const asciicow string = `
           (    )
            (oo)
   )\.-----/(O O)
  # ;       / u
    (  .   |} )
     |/ ".;|/;
     "     " "
`

func (c *cow) say(w http.ResponseWriter, r *http.Request) {

	msg := "Moooo"
	c.replies++
	ua := r.UserAgent()
	log.Printf("%v requested: %v host: %v, user-agent: %s", c.replies, r.RequestURI, r.RemoteAddr, ua)
	// w.Write([]byte(msg))
	fmt.Fprintln(w, msg, asciicow)

}
