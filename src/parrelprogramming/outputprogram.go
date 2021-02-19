package parrelprogramming

import (
	"fmt"
	"sync"
)

var (
	wg = sync.WaitGroup{}
)

func main() {
	arr := []int{1, 2, 3}
	alphaarr := []byte{'a', 'b', 'c'}
	chan12 := make(chan bool, 1)
	chan21 := make(chan bool, 1)

	go printNum(arr, chan12, chan21)
	go printAlpha(alphaarr, chan12, chan21)
	chan12 <- true
	wg.Wait()
}

func printNum(arr []int, c12, c21 chan bool) {
	wg.Add(1)
	for _, item := range arr {
		<-c12
		fmt.Println(item)
		c21 <- true
	}
	wg.Done()
}

func printAlpha(arr []byte, c12, c21 chan bool) {
	wg.Add(1)
	for _, item := range arr {
		<-c21
		fmt.Println(item)
		c12 <- true
	}
	wg.Done()
}
