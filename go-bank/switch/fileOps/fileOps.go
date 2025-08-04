package fileOps

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func GetFloatNumberFromFile(fileName string) (float64, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return 1000, errors.New("could not read file, initializing with default balance")
	}
	valueText := string(data)
	value, err := strconv.ParseFloat(valueText, 64)
	if err != nil {
		return 1000, errors.New("could not parse value, initializing with default balance")
	}
	return value, nil
}

func WriteFloatToFile(value float64, fileName string) {
	balanceText := fmt.Sprint(value)
	os.WriteFile(fileName, []byte(balanceText), 0644)
}
