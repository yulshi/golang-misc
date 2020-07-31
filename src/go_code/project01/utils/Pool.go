package utils

import "fmt"

var PoolName string
var PoolSize uint8

func init() {
	fmt.Println("pool is being initiated")
	PoolName = "Pool for Testing"
	PoolSize = 64
}
