package util

import (
	"os"
	"fmt"
	"errors"
	"math"
)

// GetEnvironmentVariable - Validate and get and environment variable value
func GetEnvironmentVariable(k string) (string, error) {
	v := os.Getenv(k)
	err := error(nil)

	if v == "" {
		message := fmt.Sprintf("%s environment variable not set.", k)
		err = errors.New(message)
	}

	return v, err
}

// AreApproximatelyEqual - Returns float equality within a threshold of 1e-3
func AreApproximatelyEqual(a float32, b float32) bool {
	return math.Abs(float64(a-b)) < 0.001
}
