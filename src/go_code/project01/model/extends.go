package model

import "fmt"

type Monkey struct {
  Name string
}

func (this *Monkey) Climb()  {
  fmt.Println(this.Name, "can climb")
}

type LittleMonkey struct {
  Monkey
}

type Flyable interface {
  Fly()
}

func (this *LittleMonkey) Fly()  {
  fmt.Println(this.Name, " can fly now")
}