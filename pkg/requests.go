package pkg

import (
	"log"
	"net/http"
	"sync"
	"time"
)

func MakeRequests(url string, timeout int, wg *sync.WaitGroup, responseTime chan<- float64, limit int) {
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
		responseTime <- float64(time.Since(startTime))
	}

}
