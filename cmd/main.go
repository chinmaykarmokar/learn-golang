package main

import (
	"cmd/cmd/controller"
	"fmt"
	"sort"
	"strings"
)

func main() {
	// Variables
	var nameOne string = "Chinmay"
	var nameTwo = "Tanmoy"
	nameThree := "Seema"

	var num1 int = 21
	var num2 = 30
	num3 := 52
	points := 2.3563

	// Print methods
	fmt.Println(nameOne, nameTwo, nameThree)
	fmt.Println(num1, num2, num3)

	fmt.Printf("My family: %v, %v, %v \n", nameOne, nameTwo, nameThree)
	fmt.Printf("My family: %q, %q, %q \n", nameOne, nameTwo, nameThree)
	fmt.Printf("Type of num1 is %T \n", num1)
	fmt.Printf("Barcelona average %0.2f points per game \n", points)

	// Arrays
	ages := [3]int{21, 30, 52}
	ages[0] = 22

	fmt.Println(ages, "The length of ages array is", len(ages))

	// Slices
	names := []string{"Chinmay", "Tanmoy", "Seema"}
	names = append(names, "Madhusudan")

	fmt.Println(names, "The length of the names slice is", len(names))

	// Slice Ranges
	range1 := names[0:3]
	fmt.Println(range1)

	// Built-in packages
	testStr := "Hello, my name is Chinmay karmokar and I am 21 years old."

	fmt.Println("Original testStr", testStr)
	fmt.Println(strings.Contains(testStr, "Chinmay"))
	fmt.Println(strings.ReplaceAll(testStr, "Hello", "Hi"))
	fmt.Println(strings.Index(testStr, "Chinmay"))
	fmt.Println(strings.ToUpper(testStr))
	fmt.Println(strings.ToLower(testStr))

	randomNumbers := []int{2, 34, 32, 54, 67, 75, 32, 21, 29, 65}

	sort.Ints(randomNumbers)
	fmt.Println(randomNumbers)
	fmt.Println(sort.SearchInts(randomNumbers, 32))

	footballTeams := []string{"Barcelona", "United", "PSG", "Chelsea", "Arsenal", "Bayern", "Dortmund"}

	sort.Strings(footballTeams)
	fmt.Println(footballTeams)
	fmt.Println(sort.SearchStrings(footballTeams, "Barcelona"))

	// Loops
	for i := 0; i < 10; i++ {
		fmt.Println("i is equal to: ", i)
	}
	for i := 0; i < len(footballTeams); i++ {
		fmt.Println(footballTeams[i])
	}

	// Booleans and Conditionals
	legalAge := 18

	fmt.Println(legalAge > 15)
	fmt.Println(legalAge < 10)

	if legalAge < 18 {
		fmt.Println("User can vote!")
	} else {
		fmt.Println("user cannot vote :(")
	}

	for index := range footballTeams {
		if index == 1 {
			continue
		} else if index == 3 {
			break
		}

		fmt.Println("Indexes are:", index)
	}

	// Functions
	additionOfNumbers := controller.AddNumbers(3, 5)
	areaOfCircle := controller.AreaOfCircle(5)

	fmt.Println("The addition of the two numbers is:", additionOfNumbers)
	fmt.Printf("The area of the circle is: %0.2f square units.\n", areaOfCircle)

	// Maps
	person := map[string]string{
		"name":  "Chinmay",
		"age":   "21",
		"hobby": "Cycling",
	}

	for key, value := range person {
		fmt.Println(key, "-", value)
	}

	fmt.Println(person)
	fmt.Println(person["hobby"])
}
