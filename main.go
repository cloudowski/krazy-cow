package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/spf13/viper"
	cow "gitlab.com/cloudowski/trapped-cow/pkg/cow"
	"gitlab.com/cloudowski/trapped-cow/pkg/shepherd"
)

var c cow.Cow
var cowconf *viper.Viper

func init() {
	c.Init()

	log.Printf("cow %s (%s version %s) initialized", c.Name, APPNAME, VERSION)

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

	http.Handle("/metrics", promhttp.Handler())
	http.HandleFunc("/", logging(c.Say))
	http.HandleFunc("/setfree", logging(c.SetFree))
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func logging(h http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		c.Requests++
		ua := r.UserAgent()

		log.Printf("%v requested: %v host: %v, user-agent: %s", c.Requests, r.RequestURI, r.RemoteAddr, ua)
		shepherd.SendStats(c.Name, fmt.Sprintf("%v %v %v", r.RequestURI, r.RemoteAddr, ua))
		h(w, r)
	}

}
