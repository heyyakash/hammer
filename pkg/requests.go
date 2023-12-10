package pkg

import (
	"log"
	"net/http"
	"sync"
	"time"
)

func MakeRequests(url string, timeout int, wg *sync.WaitGroup, responseTime chan<- time.Duration, limit int) {
	defer wg.Done()

	client := http.Client{
		Timeout: time.Duration(timeout) * time.Second,
	}

	for r := 1; r <= limit; r++ {
		startTime := time.Now()
		res, err := client.Get(url)
		if err != nil {
			log.Println("ran intp error", err)
		}
		res.Body.Close()
		duration := time.Since(startTime)
		responseTime <- duration
	}

}
