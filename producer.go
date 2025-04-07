package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"sync"
	"time"
)

const (
	baseURL       = "http://localhost:8080"
	totalRequests = 200
	concurrency   = 10
)

func Produce() {
	rand.Seed(time.Now().UnixNano())
	endpoints := []string{"add", "multiply"}

	var wg sync.WaitGroup
	semaphore := make(chan struct{}, concurrency)

	for i := 0; i < totalRequests; i++ {
		wg.Add(1)
		semaphore <- struct{}{}

		go func() {
			defer wg.Done()
			defer func() { <-semaphore }()

			a := rand.Intn(100)
			b := rand.Intn(100)
			task := endpoints[rand.Intn(len(endpoints))]

			url := fmt.Sprintf("%s/%s?a=%d&b=%d", baseURL, task, a, b)
			resp, err := http.Get(url)
			if err != nil {
				fmt.Printf("❌ Request error: %v\n", err)
				return
			}
			defer resp.Body.Close()

			fmt.Printf("✅ %s => %s\n", task, resp.Status)
		}()
	}

	wg.Wait()
	fmt.Println("🚀 Load test completed")
}
