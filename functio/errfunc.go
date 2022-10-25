package main

import (
	"fmt"
	"os"
)

func main() {
	_, err := os.Open("test.txt")

	if err != nil {
		fmt.Println("Error", err)
	} else {
		fmt.Println("found file")
	}
}
