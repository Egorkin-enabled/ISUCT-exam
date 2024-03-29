package main

import (
	"fmt"
	"strings"

	"Exam.isuct/case_n/entries"
)

func printStats(collection []entries.IEntry) {
	const row = "| %5v | %10v | %10v | %10v |\n"

	fmt.Println("Stats:")
	head := fmt.Sprintf(row, "#", "Name", "Value", "Kind")
	fmt.Printf(head)
	fmt.Println(strings.Repeat("-", len(head)-1))

	for i, v := range collection {
		fmt.Printf(row, i, v.GetName(), v.GetValue(), v.GetKind())
	}

	fmt.Println()
}

func createInstanceInformative(collection *[]entries.IEntry, name string, weight uint32) entries.IEntry {
	i, err := entries.NewEntry(name, weight)

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
	collection := make([]entries.IEntry, 0, 4)

	createInstanceInformative(&collection, "A", 130)
	createInstanceInformative(&collection, "B", 160)
	createInstanceInformative(&collection, "C", 180)
	createInstanceInformative(&collection, "D", 000)
	createInstanceInformative(&collection, "E", 280)

	printStats(collection)

	sum := entries.CalculateValueSum(collection)
	avg, err := entries.CalculateValueAvg(collection)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Weight:\n    SUM: %8v;\n    AVG: %8.2f;\n\n", sum, avg)

	fmt.Print("Trying to AVG with an empty list:\n    ")

	if avg, err := entries.CalculateValueAvg([]entries.IEntry{}); err != nil {
		fmt.Printf("Successful error: %v\n", err)
	} else {
		fmt.Printf("Unexcepted success: %v\n", avg)
	}

	fmt.Println()

	var success bool

	fmt.Print("Trying to add unique object:\n    ")
	success = entries.TryAddUniqueInstance(
		&collection,
		createInstanceInformative(nil, "D", 150))

	if success {
		fmt.Print("Success!\n\n")
	} else {
		fmt.Print("Unexcepted failure.\n\n")
	}

	fmt.Print("Trying to add non-unique object:\n    ")
	success = entries.TryAddUniqueInstance(
		&collection,
		createInstanceInformative(nil, "B", 160))

	if success {
		fmt.Print("Unexcepted success.\n\n")
	} else {
		fmt.Print("Successful failure!\n\n")
	}

	printStats(collection)
}
