package main

import "time"

func main() {
	// Background jobs
	go func() {
		for {
			AiringAnime()
			time.Sleep(time.Duration(2) * time.Second)
		}
	}()

	// Main loop
	for {
		time.Sleep(time.Duration(10) * time.Second)
	}
}
