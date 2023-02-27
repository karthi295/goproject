package db

import (
	"log"

	"github.com/gomodule/redigo/redis"
)

func Connection() redis.Conn {
	con, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		log.Fatal(err)
	}
	//defer con.Close()

	return con
}

// Invoke the command using the Do command
