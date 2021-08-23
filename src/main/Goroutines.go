package main

import (
	"fmt"
	"time"
)

//timing is key
func printingHelloWorld() {
	go say("world")
	say("hello")
}

func say(s string) {
	for i := 0; i < 1; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s, time.Now().Nanosecond())
	}
}

//splitting the work
func playingWithChannels() {
	s := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c
	sum := x + y
	fmt.Println(x, y, sum)

}

//no return needed sent back to c (channel)
func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	//send the sum to channel
	c <- sum
}

func channelBuffering() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	//uncomment and the code runs forever
	//overloading a buffered channel is blocking
	//ch <- 3
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	//trying to pull from an empty buffer leads to blocking as well
	//fmt.Println(<-ch)
}

func fibonacci2(n int, c chan int) {
	x, y := 0, 1
	//ran after the method ends
	defer close(c)

	for i := 0; i < n; i++ {
		//send current value to channel
		c <- x
		x, y = y, x+y
	}
}

//range channel for iteration
func rangeChannel() {
	//make channel and pass to function
	c := make(chan int, 10)
	go fibonacci2(cap(c), c)

	//print out each number received
	for i := range c {
		fmt.Println(i)
	}
}

func fibonacciselect(c, quit chan int) {
	x, y := 0, 1
	for {
		//waits on both cases, chooses at random if both are ready
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

//function that calls a channel multiple times to get each subsequent number
func selectFun() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			//pull from c ten times to get the 10th fibbonacci number
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacciselect(c, quit)
}

//timing is key again
func tickBoom() {
	tick := time.Tick(1000 * time.Millisecond)
	boom := time.After(5000 * time.Millisecond)
	for {
		select {
		case <-tick:
			fmt.Print("tick.")
		case <-boom:
			fmt.Print("BOOM!")
			//we are done
			return
		default:
			fmt.Print("    .")
			time.Sleep(50 * time.Millisecond)
		}
	}
}

func Goroutinesmain() {
	tickBoom()
}
