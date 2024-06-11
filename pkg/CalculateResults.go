package pkg

import (
	"fmt"
	"sync/atomic"
	"time"

	"github.com/heyyakash/hammer/models"
)

func CalculateResults(responseTime <-chan time.Duration, statistics *models.Stats) {
	var times []time.Duration
	for t := range responseTime {
		times = append(times, t)
	}
	fmt.Printf("\n\nHammer Results\n")
	fmt.Printf("Average time\t2xx Response\t4xx Response\t5xx Response\tError Count\tTotal Requests\n")
	fmt.Printf("%v\t\t%v\t\t%v\t\t%v\t\t%v\t\t%v\n\n",
		Average(times),
		atomic.LoadInt64(&statistics.Counter2xx),
		atomic.LoadInt64(&statistics.Counter4xx),
		atomic.LoadInt64(&statistics.Counter5xx),
		atomic.LoadInt64(&statistics.ErrorCount),
		atomic.LoadInt64(&statistics.Requests),
	)
}

func Average(times []time.Duration) time.Duration {
	var sum time.Duration
	if len(times) == 0 {
		return 0
	}
	for _, i := range times {
		sum += i
	}
	avg := sum / time.Duration((len(times)))
	return avg
}
