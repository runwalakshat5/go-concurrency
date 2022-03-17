package main

import (
	"fmt"
	"time"
)

var done8 = make(chan bool)
var tasks8 = make(chan int)

func produce8(n int) {
	i:=1
	for tick:=range time.Tick(time.Second){
        for ;;i++{
			fmt.Println(i,"is produced by",n,"at",time.Now())
			tasks8<-i

			
			if(i%n==0){
				i++;
				break
			}
		}
		
		if(i>100){
			fmt.Println(tick)
			break
		}
     }

	tasks8<-100
	done8 <- true
}
func consume8() {
	for {
		msg := <-tasks8
		fmt.Println(msg,"is consumed")
	}
		
	
}
func main() {
	go produce8(3)
	go produce8(5)
	go consume8()
	
	msg := <-done8
	fmt.Println(msg)
}
