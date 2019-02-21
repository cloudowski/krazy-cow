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
	mood     int
}

// happy threshold defines mimimum value of Mood parameter that determines if a cow is happy or not
const happyThreshold = 10
const asciicow string = `
           (    )
            (oo)
   )\.-----/(O O)
  # ;       / u
    (  .   |} )
     |/ ".;|/;
     "     " "
`

func NewCow() Cow {

	c := Cow{}

	if name, err := os.Hostname(); err != nil {
		log.Fatalln("Failed to get cow name (read hostname)")
	} else {
		c.Name = name
	}

	return c

}

func (c *Cow) SetMood(mood int) {
	c.mood = mood
}

func (c *Cow) GetMood() int {
	return c.mood
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

// Returns ok only if cow is happy (see happyThreshold)
func (c *Cow) Healthcheck(w http.ResponseWriter, r *http.Request) {
	if c.GetMood() >= happyThreshold {
		fmt.Fprint(w, "MooOK")
	} else {
		http.Error(w, fmt.Sprintf("I am not ok, mood level: %v", c.GetMood()), http.StatusBadRequest)
	}

}

// function that should run as goroutine to change cow mood
func (c *Cow) MoodChanger(intervalSeconds int, moodChange int) {
	log.Printf("Initializing MoodChanger - interval: %vs, change: %v\n", intervalSeconds, moodChange)
	for {
		time.Sleep(time.Duration(intervalSeconds) * time.Second)
		if c.GetMood() > 0 {
			c.SetMood(c.GetMood() + moodChange)
		}
		log.Println("MoodChanger - current mood:", c.GetMood())
	}
}
