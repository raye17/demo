package main

import (
	"demo/sss"
	"demo/sxy"
	"fmt"
)

func main() {
	fmt.Println("Hello, World!")
	fmt.Println("sxy")
}
func init() {
	sxy.Hello()
	sss.Hello()
}
