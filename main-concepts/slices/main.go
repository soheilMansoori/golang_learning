package main;
import "fmt";

// main function 
func main() {
	var numbers []int = [] int {1,2,3,4,5,6,7,8,9,10,11,12,13,14,15};

	// var mySlice = myArray[start:end] // A slice made from the array
	var mySlice = numbers[0 : 3]; // slice item of array => [index : index+1]

	fmt.Println("numbers => " , numbers);
	fmt.Println("mySlice => " , mySlice , cap(mySlice)); // capacity = 10

	// create slice 
	// slice_name := make([]type, length, capacity) 

	var languagesSlice = make([]string , 5 , 10);

	fmt.Println("languagesSlice => " , languagesSlice)

	var array = [] int {1,2,3,4,5};

	// push item to array 
	array = append(array , 6, 7);

	fmt.Println("array => " , array);


	// Create copy with only needed numbers
	neededNumbers := numbers[:len(numbers)-10];
	numbersCopy := make([]int, len(neededNumbers));
	copy(numbersCopy, neededNumbers);

	fmt.Println("numbers => " , numbers);
	fmt.Println("numbersCopy => ", numbersCopy)

};