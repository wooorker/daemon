package email

import "github.com/garyburd/redigo/redis"

func init() {

}

func connect() {
	c, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		// return nil, err
	}
}
