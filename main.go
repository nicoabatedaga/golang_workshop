package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// defer profile.Start(profile.ProfilePath(".")).Stop()

	c := make(chan int)
	var m sync.Mutex
	x, y := 0, 1

	go fibonacci(c, &x, &y, 1, &m)
	go fibonacci(c, &x, &y, 2, &m)

	go func() {
		time.Sleep(5 * time.Second)
		close(c)
	}()

	for i := range c {
		println(i)
	}

}

func fibonacci(c chan int, x, y *int, processID int, m *sync.Mutex) {
	for {
		m.Lock()
		if IsClosed(c) {
			m.Unlock()
			break
		}
		println(fmt.Sprintf("Process %d -> %d", processID, *x))
		c <- *x
		*x, *y = *y, *x+*y
		m.Unlock()
		// wait
		time.Sleep(500 * time.Millisecond)
	}
}

func IsClosed(ch <-chan int) bool {
	select {
	case <-ch:
		return true
	default:
	}

	return false
}
