package fileOperations

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func WriteFloatToFile(filename string, value float64) {
	valueText := fmt.Sprint(value)
	os.WriteFile(filename, []byte(valueText), 0644)
}

func GetStoredFloat(fileName string) (float64, error) {
	data, err := os.ReadFile(fileName)

	if err != nil {
		return 1000, errors.New("Failed to find file")
	}

	valueText := string(data)
	value, err := strconv.ParseFloat(valueText, 64)
	if err != nil {
		return 1000, errors.New("Failed to parse stored value")
	}
	return value, nil
}
