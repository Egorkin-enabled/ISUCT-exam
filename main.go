package main

import (
	"fmt"
	"strings"

	"Exam.isuct/case_2/animals"
)

func printStats(collection []animals.IAnimal) {
	const row = "| %5v | %10v | %10v | %10v |\n"

	fmt.Println("Stats:")
	head := fmt.Sprintf(row, "#", "Name", "Weight", "Kind")
	fmt.Printf(head)
	fmt.Println(strings.Repeat("-", len(head)-1))

	for i, v := range collection {
		fmt.Printf(row, i, v.GetName(), v.GetWeight(), v.GetKind())
	}

	fmt.Println()
}

func createInstanceInformative(collection *[]animals.IAnimal, name string, weight uint32) animals.IAnimal {
	i, err := animals.NewAnimal(name, weight)

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
	collection := make([]animals.IAnimal, 0, 4)

	createInstanceInformative(&collection, "A", 1)
	createInstanceInformative(&collection, "B", 4)
	createInstanceInformative(&collection, "C", 10)
	createInstanceInformative(&collection, "D", 8000000)

	printStats(collection)

	sum := animals.CalculateWeightSum(collection)
	avg, err := animals.CalculateWeightAvg(collection)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Weight:\n    SUM: %v;\n    AVG: %v\n\n", sum, avg)

	fmt.Print("Trying to AVG with an empty list:\n    ")

	if avg, err := animals.CalculateWeightAvg([]animals.IAnimal{}); err != nil {
		fmt.Printf("Successful error: %v\n", err)
	} else {
		fmt.Printf("Unexcepted success: %v\n", avg)
	}

	fmt.Println()

	var success bool

	fmt.Print("Trying to add unique object:\n    ")
	success = animals.TryAddUniqueInstance(
		&collection,
		createInstanceInformative(nil, "D", 80))

	if success {
		fmt.Print("Success!\n\n")
	} else {
		fmt.Print("Unexcepted failure.\n\n")
	}

	fmt.Print("Trying to add non-unique object:\n    ")
	success = animals.TryAddUniqueInstance(
		&collection,
		createInstanceInformative(nil, "B", 4))

	if success {
		fmt.Print("Unexcepted success.\n\n")
	} else {
		fmt.Print("Successful failure!\n\n")
	}

	printStats(collection)
}
