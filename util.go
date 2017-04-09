package main

import (
	"os"
	"fmt"
	"errors"
	"math"
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

func AreApproximatelyEqual(a float32, b float32) bool {
	return math.Abs(float64(a-b)) < 0.001
}
