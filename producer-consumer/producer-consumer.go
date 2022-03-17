package main

import 
("fmt"
 "time")

var done = make(chan bool)
var tasks = make(chan int)

func produce() {
	for i := 1; i <= 100; i++ {
		fmt.Printf("%d is produced", i)
		fmt.Println()
		tasks <- i
		time.Sleep(time.Second)
		
	}
	done <- true
}
func consume() {
	for {
		fmt.Printf("%d is consumed", <-tasks)		
		fmt.Println()
		
		
	}
}
func main() {
	go produce()
	go consume()

	msg := <-done
	fmt.Println(msg)
}
