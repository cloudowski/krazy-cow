package cow

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"time"

	"github.com/op/go-logging"
)

type Cow struct {
	Requests int
	Name     string
	mood     int
	say      string
}

// happy threshold defines mimimum value of Mood parameter that determines if a cow is happy or not
const happyThreshold = 10
const HeaderHttpTextClientKey = "X-IsHttpTextClient"
const asciicow string = `
           (    )
            (oo)
   )\.-----/(O O)
  # ;       / u
    (  .   |} )
     |/ ".;|/;
     "     " "
`

var logger *logging.Logger

type indexPage struct {
	Say      string
	Asciicow string
	Version  string
}

func NewCow() Cow {

	c := Cow{}

	if name, err := os.Hostname(); err != nil {
		logger.Fatal("Failed to get cow name (read hostname)")
	} else {
		c.Name = name
	}

	return c
}

func SetLogger(l *logging.Logger) {
	logger = l
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

func isTextRequest(r *http.Request) bool {
	return r.Header.Get(HeaderHttpTextClientKey) == "true"
}

func (c *Cow) Say(w http.ResponseWriter, r *http.Request) {

	msg := fmt.Sprintf("\"%s\"", c.GetSay())
	data := indexPage{Say: msg, Asciicow: asciicow, Version: "0.1.0-alpha"}
	if !isTextRequest(r) {
		tmpl := template.Must(template.ParseFiles("web/templates/index.html"))
		if err := tmpl.Execute(w, data); err != nil {
			logger.Errorf("Error formatting html template: %v", err)
		}
	} else {
		fmt.Fprintf(w, "%s\n %s\nver: %s\n", data.Say, data.Asciicow, data.Version)
	}

}

func (c *Cow) SetFree(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Moooooooo! (cow %s has been set free)", c.Name)
	go func() {
		time.Sleep(time.Second * 3)
		logger.Fatal("Cow has been set free!")
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
	logger.Debugf("Initializing MoodChanger - interval: %vs, change: %v\n", intervalSeconds, moodChange)
	for {
		time.Sleep(time.Duration(intervalSeconds) * time.Second)
		if c.GetMood() > 0 {
			c.SetMood(c.GetMood() + moodChange)
		}
		logger.Debugf("MoodChanger - current mood:", c.GetMood())
	}
}

// eat tuft from pasture id available
func (c *Cow) Grass(path string, interval int) {
	if f, err := os.Stat(path); err != nil {
		logger.Warning("Failed to check pasture:", err)
		return

	} else if f.IsDir() {
		logger.Debugf("Going to eat from %v, interval: %vs\n", path, interval)
	} else {
		logger.Error(err)
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
			logger.Warning("Could not find tuft:", err)
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
				logger.Debug("Eating tuft:", tuft)
				if err := os.Rename(filepath.Join(path, tuft), filepath.Join(path, fmt.Sprintf(".%v.eaten_by_%v", tuft, c.Name))); err != nil {
					logger.Fatal("Could not eat tuft:", err)
				}
				ate = true
				break

			} else {
				logger.Debug("NOT a tuft:", tuft)
			}

		}

		if ate {
			c.SetMood(c.GetMood() + 2)
			logger.Debug("Happiness increased - mood:", c.GetMood())
		} else if c.GetMood() > 0 {
			c.SetMood(c.GetMood() - 1)
			logger.Warning("No food - getting sad and angry - mood decreased:", c.GetMood())
		}
	}
}
