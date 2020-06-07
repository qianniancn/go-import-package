package utils

import (
	"fmt"
)

func init() {
	fmt.Println("已导入utils包")
}

// Add ...
func Add(a, b int) int {
	return a + b
}
