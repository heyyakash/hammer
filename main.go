package main

import (
	"flag"
	"log"
	"sync"
	"time"

	"github.com/heyyakash/hammer/models"
	"github.com/heyyakash/hammer/pkg"
)

var (
	targetURL   string
	requests    int
	concurrency int
	timeout     int
)

var wg sync.WaitGroup

// get flag values
func Init() {
	flag.StringVar(&targetURL, "url", "", "Target URl to benchmarking")
	flag.IntVar(&requests, "r", 100, "Total no. of requests to perform; default 100")
	flag.IntVar(&concurrency, "c", 1, "Total no. of goroutines to run; default 1")
	flag.IntVar(&timeout, "t", 20, "Request timeout; default 20s")
	flag.Parse()
}

func ValidateInputs() {
	if len(targetURL) == 0 {
		log.Fatal("Target URL cannot be empty")
	}
	if concurrency > requests {
		log.Fatal("No. of goroutines cannot be more than no. of requests")
	}
}

func main() {
	// initalize flags
	Init()

	// Validate Inputs
	ValidateInputs()

	// Initializing response time channel
	responseTime := make(chan time.Duration, requests)

	// Initializing Statistics
	statistics := models.NewStats()

	var requestsPerGoRoutine int = requests / concurrency

	for c := 1; c <= concurrency; c++ {
		wg.Add(1)
		go func() {
			pkg.MakeRequests(targetURL, timeout, &wg, responseTime, requestsPerGoRoutine, statistics)
		}()
	}

	go func() {
		wg.Wait()
		close(responseTime)
	}()

	pkg.CalculateResults(responseTime, statistics)

}
