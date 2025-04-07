package tasks

import (
	"fmt"
	"time"
)

func Multiply(a, b int) (int64, error) {
	result := int64(a * b)
	fmt.Printf("Multiply: %d * %d = %d\n", a, b, result)
	time.Sleep(5 * time.Second) // Simulate delay process
	return result, nil
}
