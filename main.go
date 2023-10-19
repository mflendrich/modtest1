package main

import (
		"fmt"

		"github.com/mflendrich/modtest2/add"
)

func main() {
	result := add.Add(2, 3)
	fmt.Printf("the great addition result is %d\n", result)
}
