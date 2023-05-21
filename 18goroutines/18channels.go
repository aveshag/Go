package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func xyz(ch chan int) {
	ch <- 1
	fmt.Println("sent value 1")
	ch <- 2
	fmt.Println("sent value 2")
	ch <- 3	
	fmt.Println("sent value 3")
	ch <- 4
	fmt.Println("sent value 4")
	ch <- 5
	fmt.Println("sent value 5")
	close(ch)
	wg.Done()
}
func main() {
	wg.Add(1)

	// create an unbuffered channel, for synchronous communication
	// ch := make(chan int) 
	// Buffered channel
	ch := make(chan int, 3)
	go xyz(ch)

	fmt.Println(<-ch)
	fmt.Println(<-ch) // Check commenting
	wg.Wait()


}