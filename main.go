package main

import (
	"fmt"
	"strings"

	"Exam.isuct/case_15/drinks"
)

func printStats(collection []drinks.IDrink) {
	const row = "| %5v | %10v | %10v | %10v |\n"

	fmt.Println("Stats:")
	head := fmt.Sprintf(row, "#", "Name", "Age", "Kind")
	fmt.Printf(head)
	fmt.Println(strings.Repeat("-", len(head)-1))

	for i, v := range collection {
		fmt.Printf(row, i, v.GetName(), v.GetAge(), v.GetKind())
	}

	fmt.Println()
}

func createInstanceInformative(collection *[]drinks.IDrink, name string, age uint32) drinks.IDrink {
	i, err := drinks.NewDrink(name, age)

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
	collection := make([]drinks.IDrink, 0, 4)

	createInstanceInformative(&collection, "A", 1008)
	createInstanceInformative(&collection, "B", 2010)
	createInstanceInformative(&collection, "C", 1819)
	createInstanceInformative(&collection, "D", 900)
	createInstanceInformative(&collection, "E", 2200)

	printStats(collection)

	sum := drinks.CalculateAgeSum(collection)
	avg, err := drinks.CalculateAgeAvg(collection)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Age:\n    SUM: %8v;\n    AVG: %8.2f;\n\n", sum, avg)

	fmt.Print("Trying to AVG with an empty list:\n    ")

	if avg, err := drinks.CalculateAgeAvg([]drinks.IDrink{}); err != nil {
		fmt.Printf("Successful error: %v\n", err)
	} else {
		fmt.Printf("Unexcepted success: %v\n", avg)
	}

	fmt.Println()

	var success bool

	fmt.Print("Trying to add unique object:\n    ")
	success = drinks.TryAddUniqueInstance(
		&collection,
		createInstanceInformative(nil, "D", 2001))

	if success {
		fmt.Print("Success!\n\n")
	} else {
		fmt.Print("Unexcepted failure.\n\n")
	}

	fmt.Print("Trying to add non-unique object:\n    ")
	success = drinks.TryAddUniqueInstance(
		&collection,
		createInstanceInformative(nil, "B", 2010))

	if success {
		fmt.Print("Unexcepted success.\n\n")
	} else {
		fmt.Print("Successful failure!\n\n")
	}

	printStats(collection)
}
