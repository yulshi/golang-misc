package main

import (
  "fmt"
  "github.com/garyburd/redigo/redis"
)

func main() {

  conn, err := redis.Dial("tcp", "127.0.0.1:6379")
  if err != nil {
    fmt.Println("Error while connecting: ", err)
    return
  }

  defer conn.Close()

  _, err = conn.Do("set", "test1", "abcdefg")
  if err != nil {
    fmt.Println("Error: ", err)
    return
  }

  res, err := redis.String(conn.Do("get", "test1"))
  if err != nil {
    fmt.Println("Error while executing command, ", err)
    return
  }

  fmt.Println(string(res))


}
