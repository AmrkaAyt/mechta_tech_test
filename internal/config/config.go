package config

import "flag"

type Config struct {
	NumWorkers int
	FilePath   string
}

func ParseFlags() Config {
	workers := flag.Int("workers", 4, "Number of goroutines to use for processing")
	filePath := flag.String("file", "data.json", "Path to the JSON file")
	flag.Parse()

	return Config{
		NumWorkers: *workers,
		FilePath:   *filePath,
	}
}
