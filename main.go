package main

import (
	"fmt"
	"strings"

	"Exam.isuct/case_8/sports"
)

func printStats(collection []sports.ISport) {
	const row = "| %5v | %10v | %10v | %10v |\n"

	fmt.Println("Stats:")
	head := fmt.Sprintf(row, "#", "Kind", "Age", "Country")
	fmt.Printf(head)
	fmt.Println(strings.Repeat("-", len(head)-1))

	for i, v := range collection {
		fmt.Printf(row, i, v.GetKind(), v.GetAge(), v.GetCountry())
	}

	fmt.Println()
}

func createInstanceInformative(collection *[]sports.ISport, kind string, age uint32) sports.ISport {
	i, err := sports.NewSport(kind, age)

	if err != nil {
		fmt.Printf("'%v' error: %v\n", kind, err)
		return nil
	}

	if collection != nil {
		*collection = append(*collection, i)
	}

	return i
}

func main() {

	fmt.Println("Creating instances...")
	collection := make([]sports.ISport, 0, 4)

	createInstanceInformative(&collection, "A", 1900)
	createInstanceInformative(&collection, "B", 2020)
	createInstanceInformative(&collection, "C", 1850)
	createInstanceInformative(&collection, "D", 1600)
	createInstanceInformative(&collection, "E", 1200)

	printStats(collection)

	sum := sports.CalculateAgeSum(collection)
	avg, err := sports.CalculateAgeAvg(collection)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Age:\n    SUM: %8v;\n    AVG: %8.2f;\n\n", sum, avg)

	fmt.Print("Trying to AVG with an empty list:\n    ")

	if avg, err := sports.CalculateAgeAvg([]sports.ISport{}); err != nil {
		fmt.Printf("Successful error: %v\n", err)
	} else {
		fmt.Printf("Unexcepted success: %v\n", avg)
	}

	fmt.Println()

	var success bool

	fmt.Print("Trying to add unique object:\n    ")
	success = sports.TryAddUniqueInstance(
		&collection,
		createInstanceInformative(nil, "D", 1920))

	if success {
		fmt.Print("Success!\n\n")
	} else {
		fmt.Print("Unexcepted failure.\n\n")
	}

	fmt.Print("Trying to add non-unique object:\n    ")
	success = sports.TryAddUniqueInstance(
		&collection,
		createInstanceInformative(nil, "B", 2020))

	if success {
		fmt.Print("Unexcepted success.\n\n")
	} else {
		fmt.Print("Successful failure!\n\n")
	}

	printStats(collection)
}
