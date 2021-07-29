package main

import (
	"fmt"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

func main() {
	msg1 := <-c1
	msg2 := <-c2
	wg.Add(3)
	go fun1()
	go fun2()
	go fun3()
	go fun4(c1, c2)
	wg.Wait()
}
func fun1() {
	c1 := make(chan int)
	for i := 0; i < 10; i++ {
		fmt.Println("fun1,  ->", i)
		c1 <- i
		time.Sleep(time.Duration(5 * time.Millisecond))
	}

	wg.Done()
}
func fun2() {
	c2 := make(chan int)
	for i := 0; i < 10; i++ {
		fmt.Println("fun2,  ->", i)
		c2 <- i
		time.Sleep(time.Duration(10 * time.Millisecond))
	}

	wg.Done()

}
func fun3() {
	msg1 := <-c1
	msg2 := <-c2
	defer close(msg1)
	defer close(msg2)
	fmt.Println(" data from c1  : ", msg1)
	fmt.Println(" data from c2 : ", msg2)
	wg.Done()
}
func fun4(c1 chan int, c2 chan int) {

	fmt.Println("fun4 data from c1  : ", c1)
	fmt.Println("fun4 data from c2 : ", c2)
	wg.Done()
}
