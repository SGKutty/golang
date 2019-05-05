// Sample program to show how to use the atomic package functions
// Store and Load to provide safe access to numeric types.

package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// shutdown is a flag to alert running goroutines to shutdown.
var shutdown int64

func main() {

	// Number of goroutines to use.
	const grs = 2

	// wg is used to manage concurrency.
	var wg sync.WaitGroup
	wg.Add(grs)

	// Create two goroutines.
	for i := 0; i < grs; i++ {
		go func(id int) {
			doWork(id)
			wg.Done()
		}(i)
	}

	// Give the goroutines time to run so we can see
	// the shutdown flag work.
	time.Sleep(time.Second)

	// Safely flag it is time to shutdown.
	fmt.Println("Shutdown Now")
	atomic.StoreInt64(&shutdown, 1)

	// Wait for the goroutines to finish.
	wg.Wait()
}

// doWork simulates a goroutine performing work and
// checking the shutdown flag to terminate early.
func doWork(id int) {
	for {
		fmt.Printf("Doing %d Work\n", id)
		time.Sleep(250 * time.Millisecond)

		// Do we nee to shutdown.
		if atomic.LoadInt64(&shutdown) == 1 {
			fmt.Printf("Shutting %d down\n", id)
			break
		}
	}
}
