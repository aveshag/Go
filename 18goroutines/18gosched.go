// Reference: https://stackoverflow.com/questions/13107958/what-exactly-does-runtime-gosched-do
package main

import (
	"fmt"
	// "runtime"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		// runtime.Gosched()
		fmt.Println(s)
	}
}

func main() {
	go say("world")
	say("hello")
}
