package pkg

import (
	"log"
	"time"
)

func CalculateResults(responseTime <-chan time.Duration) {
	var times []time.Duration
	for t := range responseTime {
		times = append(times, t)
	}
	log.Print("Average time : ", Average(times))
}

func Average(times []time.Duration) time.Duration {
	var sum time.Duration
	for _, i := range times {
		sum += i
	}
	avg := sum / time.Duration((len(times)))
	return avg
}
