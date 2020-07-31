package main

import (
  "fmt"
  "runtime"
  "time"
)

var c = make(chan int, 10)
var a string

func f() {
  time.Sleep(time.Second)
  a = "hello, world"
  //c <- 0
  <-c
}

func main() {
  go f()
  //<-c
  c <- 0
  print(a)
}


func pipe() {
  numChan := make(chan int, 20)
  flagChan := make(chan bool, 1)

  go func() {
    //for {
    //  println("len: ", len(numChan))
    //  num, ok := <-numChan
    //  if ok {
    //    fmt.Println(num)
    //    time.Sleep(time.Second * 5)
    //  } else {
    //    break
    //  }
    //}
    for num := range numChan {
      fmt.Println(num)
      time.Sleep(time.Second)
    }
    close(flagChan)
  }()

  go func() {
    for i := 1; i <= 50; i++ {
      fmt.Println("writing ...", i)
      numChan <- i
    }
    close(numChan)
  }()

  for {
    _, ok := <-flagChan
    if !ok {
      break
    }
  }

}

/**
  打印1到maxNum之间所有的素数
*/
func prime(maxNum int) {

  numCpu := runtime.NumCPU()

  intChan := make(chan int, 200)
  primeChan := make(chan int, 200)
  exitChan := make(chan bool, numCpu)

  // 查看指定的数字是否素数
  var isPrime = func(num int) bool {
    found := false
    for i := 2; i < num; i++ {
      if num%i == 0 {
        found = true
        break
      }
    }
    return !found
  }

  // 向intChan放入maxNum个数
  var putNumber = func(maxNum int) {
    for i := 1; i <= maxNum; i++ {
      intChan <- i
    }
    close(intChan)
  }

  // 从intChan中找出素数，并放入到primeChan中
  var calPrime = func() {
    for num := range intChan {
      if isPrime(num) {
        primeChan <- num
      }
    }
    exitChan <- true
  }

  go putNumber(maxNum)

  for i := 0; i < numCpu; i++ {
    go calPrime()
  }

  // 如果计算素数的task都完成了，就关闭primeChan
  go func() {
    for i := 0; i < numCpu; i++ {
      <-exitChan
    }
    close(primeChan)
  }()

  // 打印所有的素数
  for {
    prime, ok := <-primeChan
    if !ok {
      break
    }
    fmt.Printf("素数：%v\n", prime)
  }

}

func init() {
  fmt.Println("-------------------")
}
