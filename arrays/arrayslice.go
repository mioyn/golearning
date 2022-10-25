package main

import (
	"fmt"
	"strconv"
)

func main() {
	courses := make([]string, 5, 10)
	courses[0] = "hi"
	courses[1] = "hola"
	courses[2] = "greetings"
	courses[3] = "bello"
	courses[4] = "hello"
	fmt.Printf("lent %d and capacity %d\n", len(courses), cap(courses))
	fmt.Println(courses[4])
	fmt.Println(courses)

	sofs := courses[2:5]
	fmt.Printf("lent %d and capacity %d\n", len(sofs), cap(sofs))
	fmt.Println(sofs[2])
	fmt.Println(sofs)

	for i := 1; i < 17; i++ {
		sofs = append(sofs, strconv.Itoa(i))
		fmt.Printf("lent %d and capacity %d\n", len(sofs), cap(sofs))

	}

	dums := []string{"a", "b", "c"}
	sofs = append(sofs, dums...)
	fmt.Printf("lent %d and capacity %d\n", len(sofs), cap(sofs))
	fmt.Println(sofs)

}
