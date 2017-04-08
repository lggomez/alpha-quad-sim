package main

import (
	"os"
	"fmt"
	"errors"
)

func GetEnvironmentVariable(k string) (string, error) {
	v := os.Getenv(k)
	err := error(nil)

	if v == "" {
		message := fmt.Sprintf("%s environment variable not set.", k)
		err = errors.New(message)
	}

	return v, err
}