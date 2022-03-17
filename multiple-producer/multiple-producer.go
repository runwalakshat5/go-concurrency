package main

import (
	"fmt"
	"time"
)

var done2 = make(chan bool)
var tasks2 = make(chan int)
func produce2(n int) {
    for i := 1; i <= 10; i++ {
        fmt.Println(i,"is produced by Producer",n+1)
		
        tasks2 <- i
        time.Sleep(time.Second)
        
    }
    done2 <- true
}
func consume2() {
    for {
        msg := <-tasks2
        fmt.Println(msg,"is consumed")
    }
}
func main() {    
    for i := 0; i < 10; i++ {
        go produce2(i)
        time.Sleep(time.Second)
    }
	go consume2()
    msg:=<-done2
	fmt.Println(msg)
}