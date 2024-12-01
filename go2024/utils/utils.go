package utils

import (
	"io"
	"os"
)

func ReadInput() ([]byte, error) {
	f, err := os.Open("input.txt")
	if err != nil {
		return nil, err
	}

	data, err := io.ReadAll(f)
	if err != nil {
		return nil, err
	}

	if err := f.Close(); err != nil {
		return nil, err
	}

	return data, err
}
