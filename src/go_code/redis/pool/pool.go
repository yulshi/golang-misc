package main

import (
  "fmt"
  "github.com/garyburd/redigo/redis"
)

var pool *redis.Pool

func init() {
  pool = &redis.Pool{
    Dial: func() (redis.Conn, error) {
      return redis.Dial("tcp", "127.0.0.1:6379")
    },
    MaxIdle:     8,
    MaxActive:   0,
    IdleTimeout: 100,
  }
}

func main() {

  conn := pool.Get()

  _, err := conn.Do("hmset", "student_jimmy", "name", "Jimmy", "age", "7", "score", "98.5")
  if err != nil {
    fmt.Println("Error: ", err)
    return
  }

  jimmy, err := redis.StringMap(conn.Do("hgetall", "student_jimmy"))
  if err != nil {
    fmt.Println("Error in get: ", err)
    return
  }

  for k, v := range jimmy {
    fmt.Printf("%v: %v\n", k, v)
  }

}
