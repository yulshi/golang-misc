package main

import (
  "bufio"
  "fmt"
  "net"
  "os"
  "strings"
)

func main() {

  conn, err := net.Dial("tcp", "127.0.0.1:8888")
  if err != nil {
    fmt.Printf("Failed to connect to server: ", err)
    return
  }

  defer conn.Close()

  reader := bufio.NewReader(os.Stdin)

  for {
    line, err := reader.ReadString('\n')
    if err != nil {
      fmt.Println("Cannot read: ", err)
      break
    }

    _, err = conn.Write([]byte(line))
    if err != nil {
      fmt.Println("Failed to write to server: ", err)
      break
    }

    if strings.TrimSpace(line) == "exit" {
      break
    }
  }

}
