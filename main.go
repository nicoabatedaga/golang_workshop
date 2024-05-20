package main

import (
	"github.com/pkg/profile"
	"time"
)

func main() {
	defer profile.Start(profile.ProfilePath(".")).Stop()

	c := make(chan int)

	go fibonacci(c)
	go func() {
		time.Sleep(5 * time.Second)
		close(c)
	}()

	for i := range c {
		println(i)
	}

}

func fibonacci(c chan int) {
	x, y := 0, 1
	for {
		c <- x
		x, y = y, x+y
		// wait
		time.Sleep(500 * time.Millisecond)
	}
}
