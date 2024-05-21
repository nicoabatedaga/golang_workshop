package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	// ctx, cancel := context.WithTimeout(context.Background(), 2600*time.Millisecond)
	// ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(3*time.Second))

	result := make(chan string)

	startTime := time.Now()

	go doHttpCall(ctx, result, 1, 3)
	go doHttpCall(ctx, result, 2, 2)
	go doHttpCall(ctx, result, 3, 1)
	go doHttpCall(ctx, result, 4, 6)
	msg := <-result
	cancel()

	elapsed := time.Since(startTime)
	fmt.Println("Total time: ", elapsed)
	fmt.Println("Message: ", msg)

	time.Sleep(10 * time.Second)
	// cancel()
}

func doHttpCall(ctx context.Context, result chan string, process int, seconds int) {
	start := time.Now()
	select {
	case <-ctx.Done():
		fmt.Println(fmt.Sprintf("Process %d cancelled, elapsed time: %v", process, time.Since(start)))
	case <-time.After(time.Duration(seconds) * time.Second):
		println(fmt.Sprintf("Process %d finished, elapsed time: %v", process, time.Since(start)))
		result <- fmt.Sprintf("Finished process %d", process)
	}
}
