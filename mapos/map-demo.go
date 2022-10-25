package main

import "fmt"

func main() {
	leagueTitle := make(map[string]int)
	leagueTitle["Sunderland"] = 6
	leagueTitle["Newcastle"] = 4

	recentHead2HeadWins := map[string]int{
		"Sunderland": 6,
		"Newcastle":  0}

	fmt.Printf("Leauge titles: %v \nRecent Head to Heads: %v \n", leagueTitle, recentHead2HeadWins)

	testmap := map[string]int{
		"A": 1,
		"B": 2,
		"C": 3,
		"D": 4,
		"E": 5,
		"F": 6,
		"G": 7,
		"H": 8,
		"I": 9,
	}
	for mapKey, mapVal := range testmap {
		fmt.Printf("Key is %v Value is: %v\n", mapKey, mapVal)
	}
	fmt.Println(testmap["C"])

	testmap["A"] = 100
	fmt.Println(testmap)
	testmap["W"] = 200

	fmt.Println(testmap)

	delete(testmap, "W")
	fmt.Println(testmap)
}
