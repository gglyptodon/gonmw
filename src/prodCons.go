

package main
/* producer-consumer problem in Go */


import (
	"fmt"
	"runtime"
)

var done = make(chan bool)
var msgs = make(chan int)
func produce () {
	for i := 0; i < 10000000; i++ {
		msgs <- i
	}
	done <- true
}

func consume () {
	for {
		msg := <-msgs
		fmt.Println(msg)
	}
}
func consume2 () {
	for {
		msg := <-msgs
		fmt.Println(msg)
	}
}
func Consumer(limit int){
	for i:=0;i<limit;i++{
		go consume()
	}
}
func Producer(limit int){
	for i:=0;i<limit;i++{
		go produce()
	}
}

func main () {
	runtime.GOMAXPROCS(4)
	//go produce()
	Producer(4)
	Consumer(4)
	//go consume()
	<- done
}
