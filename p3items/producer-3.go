package main

import (
	"fmt"
	"time"
)

var done6 = make(chan bool)
var tasks6 = make(chan int)

func produce6() {
	i:=1
	for tick:=range time.Tick(time.Second){
        fmt.Printf("%d is produced", i)
		fmt.Println()
		tasks6 <- i
		i++
		fmt.Printf("%d is produced", i)
		fmt.Println()
		tasks6 <- i
		i++
		fmt.Printf("%d is produced", i)
		fmt.Println()
		tasks6 <- i
		i++
		
		if(i>=100){
			fmt.Println(tick)
			break
		}
     }

	
	done6 <- true
}
func consume6() {
	for {
		msg := <-tasks6
		fmt.Printf("%d is consumed",msg)
		fmt.Println()
	}
		
	
}
func main() {
	go produce6()

	go consume6()
	
	msg := <-done6
	fmt.Println(msg)
}
