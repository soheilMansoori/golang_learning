package main

import (
	"fmt"
	"math"
)

// Functions
func showMessage(message string) string {
	return message
}

func myFunction(x int, y int) int {
	return x + y
}

// calculator
func plus(num1 int, num2 int) int {
	return num1 + num2
}

func minus(num1 int, num2 int) int {
	return num1 - num2
}

func pow(num1 float64, num2 float64) float64 {
	result := math.Pow(num1, num2)
	return result
}

func div(num1 int, num2 int) int {
	return num1 / num2
}

func main() {
	// create variable
	var name string = "soheil" 
	job := "i am Go and Node dev" // typed declaration with initial value
	const age = 21

	// print variables
	fmt.Print(job, "\n")
	fmt.Println(name)                                              // print my name
	fmt.Printf("my var value is %v and type is %T \n", name, name) // print my variable type and value
	fmt.Printf("my age value is %v and type is %T \n", age, age)

	// create array
	//  var array_name = [length]datatype{values} // here length is defined
	var languages = [4]string{"GO", "Python", "Node", "PHP"}
	numbs := [...]int{4, 5, 6, 7, 8, 12}
	emptyIntArray := [5]int{}       //not initialized
	emptyStringArray := [5]string{} //not initialized
	emptyBooleanArray := [5]bool{}  //not initialized
	arr1 := [7]int{1: 10, 2: 40, 20, 30, 50, 60}
	// slice a array
	mySlice := arr1[2:4]

	fmt.Println(languages)
	fmt.Println(numbs)
	fmt.Println(languages[0])      // index one
	fmt.Println(emptyIntArray)     // array of 0
	fmt.Println(emptyStringArray)  // array of ""
	fmt.Println(emptyBooleanArray) // array of false
	fmt.Println(arr1)              // start form index 1
	fmt.Println(len(arr1))         // length
	fmt.Println(cap(arr1))
	// slice
	fmt.Printf("mySlice = %v\n", mySlice)
	fmt.Printf("length = %d\n", len(mySlice))
	fmt.Printf("capacity = %d\n", cap(mySlice))

	// Operators
	const num1 = 8
	const num2 = 12
	const sum = num1 + num2
	var age2 = 21
	age2 += 1

	fmt.Println("sum is =>", sum)
	fmt.Println("age2 is =>", age2)
	fmt.Println("Comparison Operators => ", sum < age2)

	// Conditionals
	const time = 20
	if time > 21 {
		fmt.Println("if scope run")
	} else if time == 20 {
		fmt.Println("else if scope run")
	} else {
		fmt.Println("else scope run")
	}

	// switch case
	day := 1
	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Saturday")
	case 7:
		fmt.Println("Sunday")
	default:
		fmt.Println("invalid value")
	}

	// Loops
	for i := 0; i <= 100; i += 10 {
		if i == 90 {
			fmt.Println("loop broken !!!")
			break
		}
		fmt.Println(i)
	}

	// Range
	// for index, value := array|slice|map {
	// 	// code to be executed for each iteration
	// }

	fruits := [3]string{"apple", "orange", "banana"}
	for idx, val := range fruits {
		fmt.Printf("%v \t%v \n", idx, val)
	}

	// Functions
	fmt.Println(myFunction(8, 12))
	var myNewMessage string = showMessage("Hi Go Dev")
	fmt.Println("my message is => ", myNewMessage)

	// my calculate
	// fmt.Println(minus(10,2))
	// fmt.Println(plus(10,2))
	// fmt.Println(div(10,5))
	// fmt.Println(pow(4,2))

}
