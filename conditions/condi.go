package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	if len, cawLen := 270, 300; len > cawLen {
		fmt.Println("len is longer thatn cawlen")
	} else if len == cawLen {
		fmt.Println("both are equal")
	} else {
		fmt.Println("cawlen is longer than len")
	}
	switch "blahBlah" {
	case "k8s":
		fmt.Println("its k8s")
	case "Kubernetes":
		fmt.Println("its Kubernetes")
	default:
		fmt.Println("no case matched")
	}

	switch tmpNum := random(); tmpNum {
	case 0, 2, 4, 6, 8:
		fmt.Println("Even number", tmpNum)
	case 1, 3, 5, 7, 9:
		fmt.Println("Odd number", tmpNum)

	}
}

func random() int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(10)
}
