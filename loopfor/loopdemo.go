package main

import (
	"fmt"
)

func main() {

	for timer := 10; timer >= 0; timer-- {
		if timer == 0 {
			fmt.Println("Boom!")
			break
		}
		fmt.Println("timer", timer)
		//time.Sleep(1 * time.Second)
	}

	listex := []string{
		"test1",
		"test2",
		"test3",
		"test4"}
breakpoint:
	for index, i := range listex {
		fmt.Println(i, index)
		break breakpoint
	}

}
