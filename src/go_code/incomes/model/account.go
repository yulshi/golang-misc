package model

import "fmt"

type Account struct {
  Name    string
  balance float64
  records []Record
}

func (acc *Account) Update(amount float64, reason string) {
  acc.balance += amount
  if acc.records == nil {
    acc.records = make([]Record, 0)
  }
  acc.records = append(acc.records, Record{
    Account:  acc,
    balance: acc.balance,
    amount:   amount,
    reason:   reason,
    isIncome: amount > 0,
  })
}

func (acc *Account) GetRecords() []Record {
  return acc.records
}

type Record struct {
  Account  *Account
  balance float64
  amount   float64
  reason   string
  isIncome bool
}

func (record *Record) String() string {
  typeText := "收入"
  if !record.isIncome {
    typeText = "支出"
  }
  return fmt.Sprintf("%v\t%.2f\t%.2f\t%v",
    typeText, record.amount, record.balance, record.reason)
}
