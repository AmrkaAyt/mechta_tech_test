package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"mechta_tech_test/internal/config"
	"mechta_tech_test/internal/processor"
	"mechta_tech_test/internal/reader"
)

func main() {
	cfg := config.ParseFlags()

	data, err := reader.ReadFile(cfg.FilePath)
	if err != nil {
		log.Fatalf("Ошибка чтения файла: %v", err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	startTime := time.Now()
	totalSum := processor.ProcessData(ctx, data, cfg.NumWorkers)

	fmt.Printf("Общая сумма: %d\n", totalSum)
	fmt.Printf("Время: %v\n", time.Since(startTime))
}
