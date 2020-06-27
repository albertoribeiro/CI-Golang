package main

import (
	"fmt"
)

func main() {
	fmt.Printf("Soma 5 + 5 = %d", soma(5, 5))
}

func soma(a int, b int) int {
	return a + b
}
