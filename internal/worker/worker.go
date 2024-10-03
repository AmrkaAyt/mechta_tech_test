package worker

import (
	"context"
	"mechta_tech_test/internal/models"
	"sync"
)

func Worker(ctx context.Context, dataChan <-chan []models.Data, resultChan chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	for chunk := range dataChan {
		sum := 0
		for _, item := range chunk {
			select {
			case <-ctx.Done():
				return
			default:
				sum += item.A + item.B
			}
		}
		resultChan <- sum
	}
}
