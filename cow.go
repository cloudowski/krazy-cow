package main

import (
	"fmt"
	"log"
	"net/http"
)

type cow struct {
	replies int
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

func (c *cow) say(w http.ResponseWriter, r *http.Request) {

	msg := fmt.Sprintf("\"%s\"", "Mooo")
	c.replies++
	ua := r.UserAgent()
	log.Printf("%v requested: %v host: %v, user-agent: %s", c.replies, r.RequestURI, r.RemoteAddr, ua)
	fmt.Fprintf(w, "%24s %s", msg, asciicow)

}
