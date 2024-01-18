package main

import (
	"fmt"
	"strings"

	"Exam.isuct/case_3/books"
)

func printStats(collection []books.IBook) {
	const row = "| %5v | %10v | %10v | %10v |\n"

	fmt.Println("Stats:")
	head := fmt.Sprintf(row, "#", "Name", "Age", "Author")
	fmt.Printf(head)
	fmt.Println(strings.Repeat("-", len(head)-1))

	for i, v := range collection {
		fmt.Printf(row, i, v.GetName(), v.GetAge(), v.GetAuthor())
	}

	fmt.Println()
}

func createInstanceInformative(collection *[]books.IBook, name string, weight uint32) books.IBook {
	i, err := books.NewBook(name, weight)

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
	collection := make([]books.IBook, 0, 4)

	createInstanceInformative(&collection, "A", 1800)
	createInstanceInformative(&collection, "B", 2020)
	createInstanceInformative(&collection, "C", 1927)
	createInstanceInformative(&collection, "D", 300)
	createInstanceInformative(&collection, "E", 2030)

	printStats(collection)

	sum := books.CalculateAgeSum(collection)
	avg, err := books.CalculateAgeAvg(collection)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Weight:\n    SUM: %8v;\n    AVG: %8.2f;\n\n", sum, avg)

	fmt.Print("Trying to AVG with an empty list:\n    ")

	if avg, err := books.CalculateAgeAvg([]books.IBook{}); err != nil {
		fmt.Printf("Successful error: %v\n", err)
	} else {
		fmt.Printf("Unexcepted success: %v\n", avg)
	}

	fmt.Println()

	var success bool

	fmt.Print("Trying to add unique object:\n    ")
	success = books.TryAddUniqueInstance(
		&collection,
		createInstanceInformative(nil, "D", 2010))

	if success {
		fmt.Print("Success!\n\n")
	} else {
		fmt.Print("Unexcepted failure.\n\n")
	}

	fmt.Print("Trying to add non-unique object:\n    ")
	success = books.TryAddUniqueInstance(
		&collection,
		createInstanceInformative(nil, "B", 2020))

	if success {
		fmt.Print("Unexcepted success.\n\n")
	} else {
		fmt.Print("Successful failure!\n\n")
	}

	printStats(collection)
}
