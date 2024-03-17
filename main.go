package main

import "sync"

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		runSMTPServer()
	}()

	go func() {
		defer wg.Done()
		runHTTPServer()
	}()

	wg.Wait()
}
