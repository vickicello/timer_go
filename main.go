package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	// Grab the first command line arg to use as the duration of the timer, like "1m30s"
	userTime := os.Args[1:][0]

	m, _ := time.ParseDuration(userTime)

	fmt.Print("This is the duration in seconds: ", time.Duration(m))
	ticker := time.NewTicker(time.Duration(m))
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			}
		}
	}()

	time.Sleep(m)
	ticker.Stop()
	done <- true
	fmt.Println("\nTimer stopped at: ", time.Now())
}
