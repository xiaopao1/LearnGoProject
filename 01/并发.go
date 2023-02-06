package main

import (
	"fmt"
	"time"
)

func hello(i int) {
	println("goroutine:" + fmt.Sprintln(i))
}
func helloGoroutine() {
	for i := 0; i < 5; i++ {
		go func(j int) {
			hello(j)
		}(i)
	}
	time.Sleep(time.Second)
}
