package main

import (
	"fmt"
	"go-unit-tests/functions"
)

func main() {
	value := functions.Sum(5, 5)
	fmt.Println(value);
}