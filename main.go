package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/spf13/viper"
	"gitlab.com/cloudowski/trapped-cow/pkg/cow"
	"gitlab.com/cloudowski/trapped-cow/pkg/shepherd"
)

var c cow.Cow
var cowconf *viper.Viper

func init() {
	c = cow.NewCow()

	log.Printf("cow %s (%s version %s) initialized", c.Name, APPNAME, VERSION)

	cowconf = viper.New()

	cowconf.SetEnvPrefix("TC")
	cowconf.AutomaticEnv()
	cowconf.SetEnvKeyReplacer(strings.NewReplacer(".", "_")) // replace "." with "_" for nested keys
	cowconf.SetConfigName("defaultconfig")                   // name of config file (without extension)
	cowconf.AddConfigPath(".")
	err := cowconf.ReadInConfig() // Find and read the config file
	if err != nil {               // Handle errors reading the config file
		log.Fatalf("Fatal error config file: %s \n", err)
	}

	cowconf.SetConfigName("cowconfig") // name of config file (without extension)
	cowconf.AddConfigPath("/config/")
	cowconf.AddConfigPath(".")
	cowconf.MergeInConfig()

	cowconf.SetDefault("cow.say", "Mooo")
	cowconf.SetDefault("logging.requests", false)
	log.Printf("Config: %v", cowconf.AllSettings())

	c.SetMood(cowconf.GetInt("cow.initmood"))

	if cowconf.GetBool("cow.moodchanger.enabled") {
		go c.MoodChanger(cowconf.GetInt("cow.moodchanger.interval"), cowconf.GetInt("cow.moodchanger.change"))
	}
}

func main() {

	http.Handle("/metrics", promhttp.Handler())
	http.HandleFunc("/", logging(c.Say))
	http.HandleFunc("/setfree", logging(c.SetFree))
	http.HandleFunc("/healthz", logging(c.Healthcheck))
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func logging(h http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		c.Requests++
		ua := r.UserAgent()

		if cowconf.GetBool("logging.requests") {
			log.Printf("%v uri: %v host: %v, user-agent: %s", c.Requests, r.RequestURI, r.RemoteAddr, ua)
		}
		shepherd.SendStats(c.Name, fmt.Sprintf("%v %v %v", r.RequestURI, r.RemoteAddr, ua))
		h(w, r)
	}

}
