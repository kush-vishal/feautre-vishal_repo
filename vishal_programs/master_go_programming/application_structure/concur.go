package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("main started")
	fmt.Println("number of cpu :", runtime.NumCPU())
	fmt.Println("number of Goroutine :", runtime.NumGoroutine())
}
