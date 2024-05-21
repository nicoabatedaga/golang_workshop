package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var times map[int]time.Duration

const calls = 4

func main() {
	times = make(map[int]time.Duration, calls)
	for i := 0; i < calls; i++ {
		times[i] = time.Duration(rand.Intn(5)) * time.Second
	}
	println(fmt.Sprintf("Times: %#v", times))

	firstApproach()
	secondApproach()
	thirdApproach()

}

func doHttpCall(i int) string {
	time.Sleep(times[i])
	return fmt.Sprintf("Ok - Call %d", i)
}

func firstApproach() {
	println("First approach")
	startTime := time.Now()

	for i := 0; i < calls; i++ {
		println(doHttpCall(i))
	}

	endTime := time.Now()

	fmt.Println("Total time first approach: ", endTime.Sub(startTime))
}

func secondApproach() {
	println("Second approach")
	startTime := time.Now()

	wg := sync.WaitGroup{}
	for i := 0; i < calls; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			println(doHttpCall(i))
		}(i)
	}
	wg.Wait()

	endTime := time.Now()

	fmt.Println("Total time second approach: ", endTime.Sub(startTime))
}

func thirdApproach() {
	println("Third approach")
	startTime := time.Now()

	result := make(chan string)
	for i := 0; i < calls; i++ {
		go func(i int) {
			result <- doHttpCall(i)
		}(i)
	}

	for i := 0; i < calls; i++ {
		println(<-result)
	}

	endTime := time.Now()

	fmt.Println("Total time third approach: ", endTime.Sub(startTime))
}
