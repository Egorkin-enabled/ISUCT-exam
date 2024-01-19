package main

import (
	"fmt"
	"strings"

	"Exam.isuct/case_1/humans"
)

func printStats(collection []humans.IHuman) {
	const row = "| %5v | %10v | %10v | %10v |\n"

	fmt.Println("Stats:")
	head := fmt.Sprintf(row, "#", "Name", "Age", "Status")
	fmt.Printf(head)
	fmt.Println(strings.Repeat("-", len(head)-1))

	for i, v := range collection {
		fmt.Printf(row, i, v.GetName(), v.GetAge(), v.GetStatus())
	}

	fmt.Println()
}

func createInstanceInformative(collection *[]humans.IHuman, name string, weight uint32) humans.IHuman {
	i, err := humans.NewHuman(name, weight)

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
	collection := make([]humans.IHuman, 0, 4)

	createInstanceInformative(&collection, "Peter", 10)
	createInstanceInformative(&collection, "Andrey", 40)
	createInstanceInformative(&collection, "Stefan", 60)
	createInstanceInformative(&collection, "Hotabich", 3000)

	printStats(collection)

	sum := humans.CalculateAgeSum(collection)
	avg, err := humans.CalculateAgeAvg(collection)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Weight:\n    SUM: %8v;\n    AVG: %8.2f;\n\n", sum, avg)

	fmt.Print("Trying to AVG with an empty list:\n    ")

	if avg, err := humans.CalculateAgeAvg([]humans.IHuman{}); err != nil {
		fmt.Printf("Successful error: %v\n", err)
	} else {
		fmt.Printf("Unexcepted success: %v\n", avg)
	}

	fmt.Println()

	var success bool

	fmt.Print("Trying to add unique object:\n    ")
	success = humans.TryAddUniqueInstance(
		&collection,
		createInstanceInformative(nil, "Jhon", 20))

	if success {
		fmt.Print("Success!\n\n")
	} else {
		fmt.Print("Unexcepted failure.\n\n")
	}

	fmt.Print("Trying to add non-unique object:\n    ")
	success = humans.TryAddUniqueInstance(
		&collection,
		createInstanceInformative(nil, "Andrey", 40))

	if success {
		fmt.Print("Unexcepted success.\n\n")
	} else {
		fmt.Print("Successful failure!\n\n")
	}

	printStats(collection)
}
