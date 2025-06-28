package main

import (
	"fmt"
	"sync"
)

func main() {
	var n int = 200000000000000000
	var num int
	mut := sync.Mutex{}
	go func() {
		for i := 0; i < n; i++ {
			mut.Lock()
			num++
			mut.Unlock()
		}
	}()
	go func() {
		for i := 0; i < n; i++ {
			mut.Lock()
			num--
			mut.Unlock()
		}
	}()

	fmt.Println(num)
}
