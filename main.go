package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
	// "github.com/heyyakash/hammer/handlers"
)

var wg sync.WaitGroup

func main() {

	url := "https://google.com"
	startTime := time.Now()
	responseTime := make(chan float64, 10)

	for i := 0; i <= 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			duration := makeRequest(url, 20)
			responseTime <- duration.Seconds()
		}()
	}
	fmt.Print("Hello")
	go func() {
		wg.Wait()
		close(responseTime)
	}()

	var durations []float64
	for t := range responseTime {
		durations = append(durations, t)
	}

	for _, t := range durations {
		fmt.Printf("%f\n", t)
	}

	log.Print("Completed in ", time.Since(startTime), "s")
	// Print results

	// mux := http.NewServeMux()

	// mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("pong"))
	// })

	// mux.HandleFunc("/start", handlers.HandleStart)

	// mux.Handle("/", http.FileServer(http.Dir("./static")))

	// log.Print("Sshh!! Server is listening")
	// if err := http.ListenAndServe(":8000", mux); err != nil {
	// 	log.Fatal("Server cannot start!")
	// }

}

func makeRequest(url string, timout int) time.Duration {
	client := http.Client{
		Timeout: time.Duration(timout) * time.Second,
	}
	startTime := time.Now()
	resp, err := client.Get(url)
	if err != nil {
		log.Print("Error sending request", err)
		return 0
	}

	defer resp.Body.Close()
	return time.Since(startTime)
}
