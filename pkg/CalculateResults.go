package pkg

import (
	"fmt"
	"sync/atomic"
	"time"

	"github.com/heyyakash/hammer/models"
)

var (
	minTime time.Duration
	maxTime time.Duration
)

func CalculateResults(responseTime <-chan time.Duration, statistics *models.Stats) {
	var times []time.Duration
	for t := range responseTime {
		times = append(times, t)
	}
	fmt.Printf("\nHammer Results\n\n")
	fmt.Println("Latency Metrics:")
	fmt.Printf("  Average Latency:   %v\n", Average(times))
	fmt.Printf("  Minimum Latency:   %v\n", minTime)
	fmt.Printf("  Maximum Latency:   %v\n", maxTime)

	fmt.Println("\nResponse Metrics:")
	fmt.Printf("  2xx Response:      %v\n", atomic.LoadInt64(&statistics.Counter2xx))
	fmt.Printf("  4xx Response:      %v\n", atomic.LoadInt64(&statistics.Counter4xx))
	fmt.Printf("  5xx Response:      %v\n", atomic.LoadInt64(&statistics.Counter5xx))

	fmt.Println("\nError Metrics:")
	fmt.Printf("  Error Count:       %v\n", atomic.LoadInt64(&statistics.ErrorCount))

	fmt.Println("\nTotal Metrics:")
	fmt.Printf("  Total Requests:    %v\n", atomic.LoadInt64(&statistics.Requests))
	fmt.Println()

}

func Average(times []time.Duration) time.Duration {
	var sum time.Duration
	if len(times) == 0 {
		return 0
	}
	for _, i := range times {
		sum += i
		if i < minTime {
			minTime = i
		}
		if i > maxTime {
			maxTime = i
		}
	}
	avg := sum / time.Duration((len(times)))
	return avg
}
