package main
import (
	"fmt"
	"runtime"
	"strconv"
)

func WorkParallel(work []Work, resq chan string) {
	queue := make(chan *Work)

	ncpu := runtime.NumCPU()
	if len(work) < ncpu {
		ncpu = len(work)
	}
	runtime.GOMAXPROCS(ncpu)

	// spawn workers
	for i := 0; i < ncpu; i++ {
		go Worker(i, queue,resq)
	}

	// master: give work
	for i, item := range(work) {
		fmt.Printf("master: give work %v\n", item)
		queue <- &work[i]  // be sure not to pass &item !!!
	}

	// all work is done
	// signal workers there is no more work
	for n := 0; n < ncpu; n++ {
		queue <- nil
	}

}

func Worker(id int, queue chan *Work, resq chan string) {
	var wp *Work
	for {
		// get work item (pointer) from the queue
		wp = <-queue
		if wp == nil {
			break
		}
		fmt.Printf("worker #%d: item %v\n", id, *wp)

		go handleWorkItem(wp,resq)
	}
}

func handleWorkItem(w *Work, resq chan string){

	for i:=0; i < 10000; i++{
		//fmt.Println("bla",w.id,i)
		if w==nil{
			resq <- "done"
		}else{
			resq <- "bla"+strconv.Itoa(w.id)
		}

	}

}
func WorkReader(resq chan string){
	var res string
	for {
		res = <-resq
		if res =="done"{
			break
		}else{fmt.Println(res)}

	}
}
type Work struct{
	id int
}
func main(){
	var lottaWork []Work
	resq := make(chan string)
	go WorkReader(resq)
	for i:=0; i< 100; i++{
		lottaWork=append(lottaWork, Work{id:i})
	}
	WorkParallel(lottaWork,resq)

	//res := resq
	//ncpu := runtime.NumCPU()
	//for n := 0; n < ncpu; n++ {
//		resq <- "done"
//	}
	//var w string
	//for {
	//	w = <-resq
	//	if w == "done"{break}
	//	fmt.Println(w)
	//}
	//for v := range resq{
	//	fmt.Println(v)
	//}
}
