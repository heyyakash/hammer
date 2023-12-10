package main

import (
	"flag"
	"sync"
	"time"

	"github.com/heyyakash/hammer/pkg"
)

var (
	targetURL   string
	requests    int
	concurrency int
	timeout     int
)

var wg sync.WaitGroup

func main() {
	flag.StringVar(&targetURL, "url", "http://localhost:8000", "Target URl to benchmarking")
	flag.IntVar(&requests, "r", 100, "Total no. of requests to perform")
	flag.IntVar(&concurrency, "c", 1, "Total no. of goroutines to run")
	flag.IntVar(&timeout, "t", 20, "Request timeout")
	flag.Parse()

	responseTime := make(chan time.Duration, requests)
	var requestsPerGoRoutine int = requests / concurrency

	for c := 1; c <= concurrency; c++ {
		wg.Add(1)
		go func() {
			pkg.MakeRequests(targetURL, timeout, &wg, responseTime, requestsPerGoRoutine)
		}()
	}

	go func() {
		wg.Wait()
		close(responseTime)
	}()

	pkg.CalculateResults(responseTime)

}
