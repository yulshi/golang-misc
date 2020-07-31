package main

import (
  "encoding/binary"
  "encoding/json"
  "errors"
  "fmt"
  common "go_code/messaging/common/message"
  "net"
)

func login(userId int, userPwd string) (err error) {

  conn, err := net.Dial("tcp", "127.0.0.1:8889")
  if err != nil {
    fmt.Println("Dial失败，", err)
    return
  }

  defer conn.Close()

  var message common.Message
  message.Type = common.Type_Login

  // 创建LoginMessage
  loginMessage := common.LoginMessage{
    UserId:   userId,
    UserPwd:  userPwd,
    UserName: "",
  }

  data, err := json.Marshal(loginMessage)
  if err != nil {
    fmt.Println("JSON序列化失败:", err)
    return err
  }

  message.Data = string(data)

  data, err = json.Marshal(message)
  if err != nil {
    fmt.Println("JSON序列化失败:", err)
    return err
  }

  // 计算长度并写成字节序列
  dataLen := uint32(len(data))
  var bytes [4]byte
  binary.BigEndian.PutUint32(bytes[:], dataLen)

  n, err := conn.Write(bytes[:])
  if err != nil {
    fmt.Println("Failed to send lenth,", err)
    return nil
  }

  if n == 0 {
    return errors.New("no data is written to server")
  }

  _, err = conn.Write(data)
  if err != nil {
   fmt.Println("Failed to send data,", err)
   return nil
  }

  return nil

}
