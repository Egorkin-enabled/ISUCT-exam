package main

import (
	"fmt"

	"Exam.isuct/case_1/humans"
)

func tryInitialize(people *[]humans.IHuman, age int32, name string) bool {
	human, err := humans.NewHuman(age, name)

	if err != nil {
		fmt.Printf("'%v' error! Message:", name)
		fmt.Println(err)
		return false
	}

	*people = append(*people, human)
	return true
}

func printStats(people []humans.IHuman) {
	fmt.Println("People stats:")
	for _, s := range people {
		fmt.Printf("name: %v; age: %v; status: %v\n", s.GetAge(), s.GetName(), s.GetStatus())
	}
}

func main() {
	people := make([]humans.IHuman, 0, 4)

	tryInitialize(&people, 10, "Jhon")
	tryInitialize(&people, 32, "Smith")
	tryInitialize(&people, 120, "Salamon")
	tryInitialize(&people, 3000, "Hatabich")

	printStats(people)

	sum := humans.CalculateAgeSum(people)
	fmt.Printf("Age sum is: %v\n", sum)

	fmt.Println("TryAddUnique - case 'failure':")

	if human, err := humans.NewHuman(32, "Smith"); err != nil {
		panic(err)
	} else if humans.TryAddUnique(&people, human) {
		fmt.Println("Unexcepted success.")
	} else {
		fmt.Println("Failure.")
	}

	printStats(people)

	fmt.Println("TryAddUnique - case 'success':")

	if human, err := humans.NewHuman(12, "Stefan"); err != nil {
		panic(err)
	} else if humans.TryAddUnique(&people, human) {
		fmt.Println("Success.")
	} else {
		fmt.Println("Unexcepted failure.")
	}

	printStats(people)
}
