package main

import (
	"fmt"
	"time"
)

var done4 = make(chan bool)
var tasks4 = make(chan int)

func produce4() {
i:=1
for{
	loop:
	for timeout := time.After(2*time.Second);;{
		select {
		case <-timeout:
			time.Sleep(5*time.Second)
			break loop
		default:		
		}
		i++
		tasks4<-i
		fmt.Printf("%d is produced", i)
		fmt.Println()
		if(i>=1000000){
			break
		}
	}
	if(i>=1000000){
		break
	}
	
}
	done4 <- true
}
func consume4() {
	for {
		fmt.Printf("%d is consumed", <-tasks4)		
		fmt.Println()
	}
}
func main() {
	go produce4()

	go consume4()
	
	msg := <-done4
	fmt.Println(msg)
}
