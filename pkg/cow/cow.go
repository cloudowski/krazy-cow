package cow

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"time"
)

type Cow struct {
	Requests int
	Name     string
	mood     int
	say      string
}

// var colors map[string]Color = {
// 	green: Color.New(FgGreen).SprintFunc(),
// }

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

func (c *Cow) SetSay(say string) {
	c.say = say
}

func (c *Cow) GetSay() string {
	return c.say
}

func (c *Cow) Say(w http.ResponseWriter, r *http.Request) {

	msg := fmt.Sprintf("\"%s\"", c.GetSay())
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
		fmt.Fprintln(w, "MooOK")
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

// eat tuft from pasture id available
func (c *Cow) Grass(path string, interval int) {
	if f, err := os.Stat(path); err != nil {
		log.Println("Failed to check pasture:", err)
		return

	} else if f.IsDir() {
		log.Printf("Going to eat from %v, interval: %vs\n", path, interval)
	} else {
		log.Println(err)
		return
	}

	tuftGreen := regexp.MustCompile("^tuft")
	tuftEaten := regexp.MustCompile("\\.eaten_by")

	var ate bool

	for {
		time.Sleep(time.Duration(interval) * time.Second)

		ate = false

		files, err := ioutil.ReadDir(path)

		if err != nil {
			log.Println("Could not find tuft:", err)
			continue
		}

		for _, f := range files {
			if f.IsDir() {
				continue
			}
			tuft := f.Name()
			if tuftEaten.MatchString(tuft) {

				continue
			} else if tuftGreen.MatchString(tuft) {
				log.Println("Eating tuft:", tuft)
				if err := os.Rename(filepath.Join(path, tuft), filepath.Join(path, fmt.Sprintf(".%v.eaten_by_%v", tuft, c.Name))); err != nil {
					log.Fatalln("Could not eat tuft:", err)
				}
				ate = true
				break

			} else {
				log.Println("NOT a tuft:", tuft)
			}

		}

		if ate {
			c.SetMood(c.GetMood() + 2)
			log.Println("Happiness increased - mood:", c.GetMood())
		} else if c.GetMood() > 0 {
			c.SetMood(c.GetMood() - 1)
			log.Println("No food - getting sad and angry - mood decreased:", c.GetMood())
		}
	}
}
