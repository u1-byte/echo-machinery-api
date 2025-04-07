package tasks

import (
	"fmt"
	"time"
)

func Add(a, b int) (int64, error) {
	result := int64(a + b)
	fmt.Printf("Add: %d + %d = %d\n", a, b, result)
	time.Sleep(2 * time.Second) // Simulate delay process
	return result, nil
}
