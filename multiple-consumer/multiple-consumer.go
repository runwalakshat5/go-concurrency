package main

import("fmt"
"time")

var done1 = make(chan bool)
var tasks1 = make(chan int)
func produce1() {
    for i := 1; i <= 1000; i++ {
        fmt.Printf("%d is produced", i)
		fmt.Println()
        tasks1 <- i
        time.Sleep(time.Second)
    }
    done1 <- true
}
func consume1(n int) {
    for {
        msg := <-tasks1
        fmt.Printf("%d is consumed by Consumer - %d ",msg,n+1)
        fmt.Println()
    }
}
func main() {
    go produce1()
    for i := 0; i < 10; i++ {
        go consume1(i)
    }
    msg:=<-done1
	fmt.Println(msg)
}