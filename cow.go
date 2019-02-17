package main

import (
	"fmt"
	"net/http"
)

type cow struct {
	requests int
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
	fmt.Fprintf(w, "%24s %s", msg, asciicow)

}
