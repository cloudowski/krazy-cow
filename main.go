package main

import (
	"log"
	"net/http"
	"strings"

	"github.com/spf13/viper"
)

var c cow
var cowconf *viper.Viper

func init() {
	c.init()

	log.Printf("cow %s (%s version %s) initialized", c.name, APPNAME, VERSION)

	cowconf = viper.New()

	cowconf.SetEnvPrefix("TC")
	cowconf.AutomaticEnv()
	cowconf.SetEnvKeyReplacer(strings.NewReplacer(".", "_")) // replace "." with "_" for nested keys
	cowconf.SetConfigName("cowconfig")                       // name of config file (without extension)
	cowconf.AddConfigPath(".")                               // optionally look for config in the working directory
	err := cowconf.ReadInConfig()                            // Find and read the config file
	if err != nil {                                          // Handle errors reading the config file
		log.Fatalf("Fatal error config file: %s \n", err)
	}

	cowconf.SetDefault("cow.say", "Mooo")
	log.Printf("Config: %v", cowconf.Get("cow.say"))
}

func main() {

	http.HandleFunc("/", logging(c.say))
	http.HandleFunc("/setfree", logging(c.setfree))
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
