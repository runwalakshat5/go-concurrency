package main

import (
	"fmt"
	"time"
)

var done5 = make(chan bool)
var tasks5 = make(chan int)

func produce5() {
	for i := 1; i <= 1000000; i++ {
		fmt.Printf("%d is produced", i)
		fmt.Println()
        tasks5 <- i
    }

	done5 <- true
}
func consume5() {
	for {
		loop:
		for timeout := time.After(2*time.Second);;{
		select {
		case <-timeout:
			time.Sleep(5*time.Second)
			break loop
		default:		
		}
		msg := <-tasks5
		fmt.Printf("%d is consumed", msg)
		fmt.Println()
	}
		
	}
}
func main() {
	go produce5()

	go consume5()
	
	msg := <-done5
	fmt.Println(msg)
}
