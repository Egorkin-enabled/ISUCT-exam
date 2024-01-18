package main

import (
	"fmt"
	"strings"

	"Exam.isuct/case_20/buildings"
)

func printStats(collection []buildings.IBuilding) {
	const row = "| %5v | %10v | %10v | %10v |\n"

	fmt.Println("Stats:")
	head := fmt.Sprintf(row, "#", "Model", "Age", "Class")
	fmt.Printf(head)
	fmt.Println(strings.Repeat("-", len(head)-1))

	for i, v := range collection {
		fmt.Printf(row, i, v.GetAddress(), v.GetAge(), v.GetClass())
	}

	fmt.Println()
}

func createInstanceInformative(collection *[]buildings.IBuilding, name string, age uint32) buildings.IBuilding {
	i, err := buildings.NewBuilding(name, age)

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
	collection := make([]buildings.IBuilding, 0, 4)

	createInstanceInformative(&collection, "A", 2008)
	createInstanceInformative(&collection, "B", 2020)
	createInstanceInformative(&collection, "D", 1600)
	createInstanceInformative(&collection, "C", 2019)
	createInstanceInformative(&collection, "E", 2200)

	printStats(collection)

	sum := buildings.CalculateAgeSum(collection)
	avg, err := buildings.CalculateAgeAvg(collection)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Age:\n    SUM: %8v;\n    AVG: %8.2f;\n\n", sum, avg)

	fmt.Print("Trying to AVG with an empty list:\n    ")

	if avg, err := buildings.CalculateAgeAvg([]buildings.IBuilding{}); err != nil {
		fmt.Printf("Successful error: %v\n", err)
	} else {
		fmt.Printf("Unexcepted success: %v\n", avg)
	}

	fmt.Println()

	var success bool

	fmt.Print("Trying to add unique object:\n    ")
	success = buildings.TryAddUniqueInstance(
		&collection,
		createInstanceInformative(nil, "D", 2001))

	if success {
		fmt.Print("Success!\n\n")
	} else {
		fmt.Print("Unexcepted failure.\n\n")
	}

	fmt.Print("Trying to add non-unique object:\n    ")
	success = buildings.TryAddUniqueInstance(
		&collection,
		createInstanceInformative(nil, "B", 2020))

	if success {
		fmt.Print("Unexcepted success.\n\n")
	} else {
		fmt.Print("Successful failure!\n\n")
	}

	printStats(collection)
}
