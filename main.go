package main

import (
	"fmt"
	"strings"

	"Exam.isuct/case_9/games"
)

func printStats(collection []games.IGame) {
	const row = "| %5v | %10v | %10v | %10v |\n"

	fmt.Println("Stats:")
	head := fmt.Sprintf(row, "#", "Name", "Age", "Genere")
	fmt.Printf(head)
	fmt.Println(strings.Repeat("-", len(head)-1))

	for i, v := range collection {
		fmt.Printf(row, i, v.GetName(), v.GetAge(), v.GetGenere())
	}

	fmt.Println()
}

func createInstanceInformative(collection *[]games.IGame, name string, age uint32) games.IGame {
	i, err := games.NewGame(name, age)

	if err != nil {
		fmt.Printf("'%v' error: %v\n", name, err)
		return nil
	}

	if collection != nil {
		*collection = append(*collection, i)
	}

	return i
}

func main() {

	fmt.Println("Creating instances...")
	collection := make([]games.IGame, 0, 4)

	createInstanceInformative(&collection, "A", 1980)
	createInstanceInformative(&collection, "B", 2020)
	createInstanceInformative(&collection, "C", 1950)
	createInstanceInformative(&collection, "D", 1600)
	createInstanceInformative(&collection, "E", 2200)

	printStats(collection)

	sum := games.CalculateAgeSum(collection)
	avg, err := games.CalculateAgeAvg(collection)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Age:\n    SUM: %8v;\n    AVG: %8.2f;\n\n", sum, avg)

	fmt.Print("Trying to AVG with an empty list:\n    ")

	if avg, err := games.CalculateAgeAvg([]games.IGame{}); err != nil {
		fmt.Printf("Successful error: %v\n", err)
	} else {
		fmt.Printf("Unexcepted success: %v\n", avg)
	}

	fmt.Println()

	var success bool

	fmt.Print("Trying to add unique object:\n    ")
	success = games.TryAddUniqueInstance(
		&collection,
		createInstanceInformative(nil, "D", 1980))

	if success {
		fmt.Print("Success!\n\n")
	} else {
		fmt.Print("Unexcepted failure.\n\n")
	}

	fmt.Print("Trying to add non-unique object:\n    ")
	success = games.TryAddUniqueInstance(
		&collection,
		createInstanceInformative(nil, "B", 2020))

	if success {
		fmt.Print("Unexcepted success.\n\n")
	} else {
		fmt.Print("Successful failure!\n\n")
	}

	printStats(collection)
}
