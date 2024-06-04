package service

import (
	"net/http"
	"sync"
	"time"

	"github.com/flavioamaral-dev/go-experts-desafio-stress-test/internal/entity"
)

type carregarTesteService struct{}

func NewLoadTesterService() entity.CarregarTeste {
	return &carregarTesteService{}
}

func (s *carregarTesteService) ExecutarTeste(config entity.TestConfig) entity.TestResult {
	var wg sync.WaitGroup
	client := &http.Client{}
	results := make(chan int, config.TotalRequests)
	start := time.Now()

	requestsPerWorker := config.TotalRequests / config.Concurrency
	for i := 0; i < config.Concurrency; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < requestsPerWorker; j++ {
				resp, err := client.Get(config.URL)
				if err != nil {
					results <- 0
					continue
				}
				results <- resp.StatusCode
				resp.Body.Close()
			}
		}()
	}

	wg.Wait()
	close(results)

	duration := time.Since(start)
	statusCounts := make(map[int]int)

	for result := range results {
		statusCounts[result]++
	}

	return entity.TestResult{
		Duration:      duration,
		TotalRequests: config.TotalRequests,
		StatusCounts:  statusCounts,
	}
}
