package shepherd

import (
	"log"

	"github.com/go-redis/redis"
)

var rc *redis.Client
var rcready bool = false

type CowStats struct {
	Requests int
	Uri      string
	Addr     string
	Agent    string
}

func init() {
	// redis

	rc = redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := rc.Ping().Result()
	if err == nil {
		log.Println("Connection to shepherd (redis) established:", pong)
		rcready = true
	} else {
		log.Printf("Error in connecting to redis: '%v'", err)
	}
}

func SendStats(id string, data string) {

	if rcready {
		if err := rc.RPush(id, data).Err(); err != nil {
			log.Println("Error sending stats to shepherd: %v", err)
		}
	}
}
