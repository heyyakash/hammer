package pkg

import (
	"net/http"
	"sync"
	"sync/atomic"
	"time"

	"github.com/heyyakash/hammer/models"
)

func MakeRequests(url string, timeout int, wg *sync.WaitGroup, responseTime chan<- time.Duration, limit int, statistics *models.Stats) {
	defer wg.Done()

	client := http.Client{
		Timeout: time.Duration(timeout) * time.Second,
	}

	for r := 1; r <= limit; r++ {
		// initializing clock
		startTime := time.Now()

		//sending HTTP GET requests
		res, err := client.Get(url)
		if err != nil {
			atomic.AddInt64(&statistics.ErrorCount, 1)
			atomic.AddInt64(&statistics.Requests, 1)
			continue
		}
		defer res.Body.Close()

		//updating counters
		switch {
		case res.StatusCode >= 200 && res.StatusCode < 300:
			atomic.AddInt64(&statistics.Counter2xx, 1)
		case res.StatusCode >= 400 && res.StatusCode < 500:
			atomic.AddInt64(&statistics.Counter4xx, 1)
		case res.StatusCode >= 500 && res.StatusCode < 600:
			atomic.AddInt64(&statistics.Counter5xx, 1)
		}

		//Incrementing request count
		atomic.AddInt64(&statistics.Requests, 1)

		// Pushing duration to channel
		duration := time.Since(startTime)
		responseTime <- duration
	}

}
