package main

import "fmt"

func main() {
	type courseMeta struct {
		author string
		level  string
		rating float64
	}

	gettingStartedwithK8s := courseMeta{
		level:  "Intermediate",
		author: "Nigel poulton",
		rating: 5,
	}

	fmt.Println(gettingStartedwithK8s)
	fmt.Println(gettingStartedwithK8s.author)

	gettingStartedwithK8s.rating = 1.3

	fmt.Println(gettingStartedwithK8s)

}
