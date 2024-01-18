package main

import (
	"fmt"
	"strings"

	"Exam.isuct/case_6/products"
)

func printStats(collection []products.IProduct) {
	const row = "| %5v | %10v | %10v | %10v |\n"

	fmt.Println("Stats:")
	head := fmt.Sprintf(row, "#", "Name", "Count", "Price")
	fmt.Printf(head)
	fmt.Println(strings.Repeat("-", len(head)-1))

	for i, v := range collection {
		fmt.Printf(row, i, v.GetName(), v.GetCount(), v.GetPrice())
	}

	fmt.Println()
}

func createInstanceInformative(collection *[]products.IProduct, name string, count uint32) products.IProduct {
	i, err := products.NewProduct(name, count)

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
	collection := make([]products.IProduct, 0, 4)

	createInstanceInformative(&collection, "A", 500)
	createInstanceInformative(&collection, "B", 600)
	createInstanceInformative(&collection, "C", 700)
	createInstanceInformative(&collection, "D", 1200)

	printStats(collection)

	sum := products.CalculatePriceSum(collection)
	avg, err := products.CalculatePriceAvg(collection)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Age:\n    SUM: %8v;\n    AVG: %8.2f;\n\n", sum, avg)

	fmt.Print("Trying to AVG with an empty list:\n    ")

	if avg, err := products.CalculatePriceAvg([]products.IProduct{}); err != nil {
		fmt.Printf("Successful error: %v\n", err)
	} else {
		fmt.Printf("Unexcepted success: %v\n", avg)
	}

	fmt.Println()

	var success bool

	fmt.Print("Trying to add unique object:\n    ")
	success = products.TryAddUniqueInstance(
		&collection,
		createInstanceInformative(nil, "D", 120))

	if success {
		fmt.Print("Success!\n\n")
	} else {
		fmt.Print("Unexcepted failure.\n\n")
	}

	fmt.Print("Trying to add non-unique object:\n    ")
	success = products.TryAddUniqueInstance(
		&collection,
		createInstanceInformative(nil, "B", 600))

	if success {
		fmt.Print("Unexcepted success.\n\n")
	} else {
		fmt.Print("Successful failure!\n\n")
	}

	printStats(collection)
}
