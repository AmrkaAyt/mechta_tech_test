package reader

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"mechta_tech_test/internal/models"
	"os"
)

func ReadFile(filePath string) ([]models.Data, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("не удалось открыть файл: %w", err)
	}
	defer file.Close()

	byteValue, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("не удалось прочитать файл: %w", err)
	}

	var data []models.Data
	if err := json.Unmarshal(byteValue, &data); err != nil {
		return nil, fmt.Errorf("ошибка unmarshal JSON: %w", err)
	}

	return data, nil
}
