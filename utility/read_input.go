package utility

import (
	"errors"
	"os"
)

func ReadInput(path string) (string, error) {
	input, err := os.ReadFile(path)
	if err != nil {
		return "", errors.New("error while reading the file")
	}
	return string(input), nil
}
