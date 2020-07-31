package main

import (
  "fmt"
  "go_code/incomes/model"
  "strconv"
)

func main() {

  account := &model.Account{
    Name: "Tom",
  }

  key := ""
  quit := false

  for {
    fmt.Println("\n----------家庭收支记账软件------------")
    fmt.Println("----------1. 收支明细 ------------")
    fmt.Println("----------2. 登记收入 ------------")
    fmt.Println("----------3. 登记支出 ------------")
    fmt.Println("----------4. 退出软件 ------------")
    fmt.Println("请选择（1-4）：")
    fmt.Scan(&key)

    switch key {
    case "1":
      fmt.Println("收支\t收支金额\t账户余额\t收支说明")
      for _, ah := range account.GetRecords() {
        fmt.Println(&ah)
      }
    case "2":
      amountAsText := ""
      fmt.Print("本次收入金额：")
      fmt.Scan(&amountAsText)
      amount, err := strconv.ParseFloat(amountAsText, 64)
      if err != nil {
        fmt.Println("您输入的金额不是数字，请重新输入")
        break
      }
      fmt.Print("本次收入说明：")
      reason := ""
      fmt.Scan(&reason)
      account.Update(amount, reason)
    case "3":
      amountAsText := ""
      fmt.Print("本次支出金额：")
      fmt.Scan(&amountAsText)
      amount, err := strconv.ParseFloat(amountAsText, 64)
      if err != nil {
        fmt.Println("您输入的金额不是数字，请重新输入")
        break
      }
      fmt.Print("本次支出说明：")
      reason := ""
      fmt.Scan(&reason)
      account.Update(-amount, reason)
    case "4":
      input := ""
      for {
        fmt.Println("Are you sure to quit? (y/n): ")
        fmt.Scan(&input)
        if input == "y" {
          quit = true
          break
        } else if input == "n" {
          break
        } else {
          fmt.Println("Your input must be either 'y' or 'n': ")
        }
      }
    }

    if quit {
      fmt.Println("您已经退出记账软件，👋")
      break
    }
  }

}
