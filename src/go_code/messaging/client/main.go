package main

import "fmt"

var userId int
var userPwd string

func main() {

  var key int
  var loop = true

  for {
    fmt.Println("--------欢迎登录多人聊天系统--------")
    fmt.Println("\t\t\t1.登录聊天室")
    fmt.Println("\t\t\t2.注册用户")
    fmt.Println("\t\t\t3.退出系统")
    fmt.Print("请选择（1-3）：")

    fmt.Scanf("%d\n", &key)

    switch key {
    case 1:
      fmt.Println("登录聊天室")
      prepareLogin()
      loop = false
    case 2:
      fmt.Println("注册用户")
      loop = false
    case 3:
      exit(&loop)
    default:
      fmt.Println("您的输入有误，请重新输入")
    }

    if !loop {
      return
    }

  }

}

func prepareLogin() {

  fmt.Print("请输入用户ID：")
  fmt.Scanf("%d\n", &userId)
  fmt.Print("请输入用户密码：")
  fmt.Scanf("%s\n", &userPwd)

  login(userId, userPwd)

}

func exit(loop *bool) {

  var yesOrNo string
  for {
    fmt.Print("are you sure to exit?(y/n): ")
    fmt.Scanln(&yesOrNo)
    switch yesOrNo {
    case "n":
      *loop = true
      return
    case "y":
      *loop = false
      return
    default:
      // ...
    }
  }
}
