package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	done := make(chan struct{})
	sqch := make(chan int)
	eventosquar := make(chan int, 15)
	oddtomerg := make(chan int)
	mertopri := make(chan int)
	go generator(ch)
	go oddevensplit(ch, eventosquar, oddtomerg)
	go squarer(eventosquar, sqch)
	go merger(oddtomerg, sqch, mertopri)
	go printer(mertopri, done)
	<-done

}
func generator(out chan int) {
	for i := 0; i < 10; i++ {
		out <- i
	}
	close(out)
}
func printer(in chan int, done chan struct{}) {
	for a := range in {

		fmt.Println(a)
	}
	done <- struct{}{}

}
func squarer(in chan int, out chan int) {
	for a := range in {
		time.Sleep(1 * time.Second)
		out <- a * a

	}
	close(out)
}
func oddevensplit(in chan int, even chan int, odd chan int) {
	for a := range in {
		if a%2 == 0 {
			even <- a
		} else {
			odd <- a
		}

	}
	close(even)
	close(odd)
}
func merger(evenin chan int, oddin chan int, out chan int) {
	oddDone := false
	evenDone := false

	for !oddDone || !evenDone {
		select {
		case a, ok := <-oddin:
			if !ok {
				oddDone = true
			} else {
				out <- a
			}
		case b, ok := <-evenin:
			if !ok {
				evenDone = true
			} else {
				out <- b
			}
		}
	}
	close(out)
}
