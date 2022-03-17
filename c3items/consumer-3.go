package main

import (
	"fmt"
	"time"
)

var done7 = make(chan bool)
var tasks7 = make(chan int)

func produce7() {
	for i := 1; i <= 100; i++ {
		fmt.Printf("%d is produced", i)
		fmt.Println()
		tasks7 <- i
	}
	
	done7 <- true
}
func consume7() {
	for tick:=range time.Tick(time.Second){
		msg := <-tasks7
		fmt.Printf("%d is consumed at %v",msg,tick)
		fmt.Println()
		msg = <-tasks7
		fmt.Printf("%d is consumed at %v",msg,tick)
		fmt.Println()
		msg = <-tasks7
		fmt.Printf("%d is consumed at %v",msg,tick)
		fmt.Println()
		
	}
		
	
}
func main() {
	go produce7()

	go consume7()
	
	msg := <-done7
	fmt.Println(msg)
}
