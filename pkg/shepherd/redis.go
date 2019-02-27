package shepherd

import (
	"fmt"
	"log"

	"github.com/go-redis/redis"
)

var rc *redis.Client
var rcready bool = false
var RedisUrl = "redis:6379"

const milkKey = "dairy"

type CowStats struct {
	Requests int
	Uri      string
	Addr     string
	Agent    string
}

func connect() error {
	// redis

	rc = redis.NewClient(&redis.Options{
		Addr:     RedisUrl,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	_, err := rc.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}

func SendStats(id string, data string) error {

	if err := connect(); err == nil {
		if err := rc.RPush(id, data).Err(); err != nil {
			return err
		}
		return nil
	} else {
		return err
	}

}

func SendMilk(id string) error {

	if err := connect(); err == nil {
		if err := rc.HIncrBy(milkKey, id, 1).Err(); err != nil {
			return fmt.Errorf("%v", err)
		}
		return nil
	} else {
		return err
	}

}

func GetStats(id string) []string {

	var result []string
	if rcready {
		llen, err := rc.LLen(id).Result()
		if err != nil {
			log.Println("Error getting stats from shepherd:", err)
			// return []string{}
			return nil
		}
		// for _,v := rc.LRange(id, 0, llen-1).Result() {
		values, err := rc.LRange(id, 0, llen-1).Result()
		if err != nil {
			log.Println("Error getting stats from shepherd:", err)
			// return []string{}
			return nil
		}
		result = append(result, values...)
		// }
		return result

	}
	return nil
}
