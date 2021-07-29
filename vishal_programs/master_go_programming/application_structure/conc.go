package main

import (
	"fmt"	"runtime"
)


func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func temp() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci(c, quit)
}
 
func main() {
    // NumCPU returns the number of logical CPUs or cores usable by the current process.
    fmt.Println("No. of CPUs:", runtime.NumCPU()) // => No. of CPUs: 4
 
    // NumGoroutine returns the number of goroutines that currently exists.
    fmt.Println("No. of Goroutines:", runtime.NumGoroutine()) // => No. of Goroutines: 1
 
    // GOOS and GOARCH are environment variables
    fmt.Println("OS:", runtime.GOOS)     // => OS: linux
    fmt.Println("Arch:", runtime.GOARCH) // => Arch: amd64
 
    //  GOMAXPROCS sets the maximum number of CPUs that can be executing simultaneously and returns
    //  the previous setting.
    fmt.Println("GOMAXPROCS:", runtime.GOMAXPROCS(0)) // => GOMAXPROCS: 4
    // If n < 1, it does not change the current setting.
 
}