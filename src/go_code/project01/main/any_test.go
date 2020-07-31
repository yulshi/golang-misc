package main

import (
  "runtime"
  "testing"
)

func TestHello(t *testing.T) {

  println(runtime.NumCPU())

}
