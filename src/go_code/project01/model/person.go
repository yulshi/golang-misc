package model

import "fmt"

type person struct {
  name   string
  age    int
  salary float64
}

func Person(name string) *person {
  return &person{
    name: name,
  }
}

func (p *person) GetName() string {
  return p.name
}

func (p *person) SetAge(age int) {
  if age < 0 || age > 200 {
    println("the age is not valid")
    return
  }
  p.age = age
}

func (p *person) GetAge() int {
  return p.age
}

func (p *person) SetSalary(salary float64) {
  if salary < 3000 {
    println("the salary is not valid")
    return
  }
  p.salary = salary
}

func (p *person) GetSalary() float64 {
  return p.salary
}

func (p *person) String() string {
  return fmt.Sprintf("person[name:%v, age: %v, salary: %.2f]", p.name, p.age, p.salary)
}
