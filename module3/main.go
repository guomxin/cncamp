package main

import (
	"fmt"
	"runtime/debug"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			debug.PrintStack()
		}
	}()

	var a any = "hello"
	b := a.(int)
	fmt.Println(b)
}
