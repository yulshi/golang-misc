package model

import "fmt"

type Student struct {
  Name  string `json:"name"`
  Age   int    `json:"age"`
  Score float64
}

func (student Student) Print() {
  fmt.Printf("Student[name=%v, age=%d, score=%.2f",
    student.Name, student.Age, student.Score)
}

func (student *Student) IncreaseScore(delta float64) {
  student.Score += delta
}

func (student *Student) Rename(newName string) string {
  oldName := student.Name
  student.Name = newName
  return oldName
}
