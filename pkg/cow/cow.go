package cow

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

type Cow struct {
	Requests int
	Name     string
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

func (c *Cow) Init() {

	if name, err := os.Hostname(); err != nil {
		log.Fatalln("Failed to get cow name (read hostname)")
	} else {
		c.Name = name
	}

}

func (c *Cow) Say(w http.ResponseWriter, r *http.Request) {

	// msg := fmt.Sprintf("\"%s\"", cowconf.Get("cow.say"))
	msg := fmt.Sprintf("\"%s\"", "Mooo - FIXME to use config")
	fmt.Fprintf(w, "%15s %s %s", " ", msg, asciicow)

}

func (c *Cow) SetFree(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Moooooooo! (cow %s has been set free)", c.Name)
	go func() {
		time.Sleep(time.Second * 3)
		log.Fatalln("Cow has been set free!")
	}()

}
