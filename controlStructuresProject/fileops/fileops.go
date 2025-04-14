package fileops

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func WriteFloatToFile(value float64, filename string) {
	valueText := fmt.Sprint(value)
	//0644 is a way of encoding file permissions. It basically means that owner can read/write file
	//other users can only read from it.
	os.WriteFile(filename, []byte(valueText), 0644)
}

func GetFloatFromFile(filename string) (float64, error) {
	//We are using the discard operator to reject the second value returned by the function
	data, err := os.ReadFile(filename)
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
