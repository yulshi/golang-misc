package main

import (
  "encoding/binary"
  "encoding/json"
  "errors"
  "fmt"
  "go_code/messaging/common/message"
  "io"
  "net"
)

func main() {

  fmt.Println("服务器在8889端口监听。。。")
  listen, err := net.Listen("tcp", "0.0.0.0:8889")
  if err != nil {
    fmt.Println("监听失败, ", err)
    return
  }

  defer listen.Close()

  for {
    fmt.Println("等待客户端连接。。。")
    conn, err := listen.Accept()
    if err != nil {
      fmt.Println("Accept失败，", err)
    } else {
      go handle(conn)
    }

  }

}

func handle(conn net.Conn) {

  defer conn.Close()

  fmt.Println("处理客户端请求：", conn.RemoteAddr().String())

  for {
    message, err := readPkg(conn)
    if err == io.EOF {
      return
    }
    if err != nil {
      fmt.Println("Failed to read package:", err)
      return
    }

    fmt.Println(message)

  }

}

func readPkg(conn net.Conn) (message message.Message, err error) {

  buff := make([]byte, 4096)
  _, err = conn.Read(buff[:4])
  if err == io.EOF {
    return
  }

  dataLen := binary.BigEndian.Uint32(buff[:4])

  // Read data
  n, err := conn.Read(buff[:dataLen])
  if err != nil {
    return
  }
  if uint32(n) != dataLen {
    err = errors.New(fmt.Sprintf("没有读到足够的数据，期望：%d，实际：%d", n, dataLen))
    return
  }

  err = json.Unmarshal(buff[:dataLen], &message)

  return

}
