package main

import (
	"fmt"
	"time"
)

func greet(phrase string, doneChan chan bool) {
	fmt.Println("Hello!", phrase)
	doneChan <- true
}

func slowGreet(phrase string, doneChan chan bool) {
	time.Sleep(3 * time.Second)
	fmt.Println("Hello!", phrase)
	doneChan <- true
	close(doneChan)
}

func main() {
	// When we have single channel for multiple go routines we have to add <- done equal to the number of go routines
	// done := make(chan bool)
	// go greet("Nice to meet you", done)
	// go greet("how are you", done)
	// go slowGreet("How are youuuuuuuuuuuu", done)
	// go greet("I hope you are doing good", done)
	// <-done
	// <-done
	// <-done
	// <-done

	// handling go routines with slice of channels
	// dones := make([]chan bool, 4)
	// dones[0] = make(chan bool)
	// go greet("Nice to meet you", dones[0])
	// dones[1] = make(chan bool)
	// go greet("how are you", dones[1])
	// dones[2] = make(chan bool)
	// go slowGreet("How are youuuuuuuuuuuu", dones[2])
	// dones[3] = make(chan bool)
	// go greet("I hope you are doing good", dones[3])

	// for _, done := range dones {
	// 	<-done
	// }

	done := make(chan bool)
	go greet("Nice to meet you", done)
	go greet("how are you", done)
	go slowGreet("How are youuuuuuuuuuuu", done)
	go greet("I hope you are doing good", done)

	for range done {
		// fmt.Println(doneChan)
	}
}
