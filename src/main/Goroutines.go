package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s, time.Now().Nanosecond())
	}
}

func routinemain() {
	go say("world")
	say("hello")
}
