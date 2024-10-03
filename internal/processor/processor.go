package processor

import (
	"context"
	"mechta_tech_test/internal/models"
	"mechta_tech_test/internal/worker"
	"sync"
)

func ProcessData(ctx context.Context, data []models.Data, numWorkers int) int {
	chunkSize := (len(data) + numWorkers - 1) / numWorkers
	dataChan := make(chan []models.Data, numWorkers)
	resultChan := make(chan int, numWorkers)

	var wg sync.WaitGroup

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker.Worker(ctx, dataChan, resultChan, &wg)
	}

	for i := 0; i < len(data); i += chunkSize {
		end := i + chunkSize
		if end > len(data) {
			end = len(data)
		}
		dataChan <- data[i:end]
	}
	close(dataChan)

	go func() {
		wg.Wait()
		close(resultChan)
	}()

	totalSum := 0
	for sum := range resultChan {
		totalSum += sum
	}

	return totalSum
}
