package main;
import "fmt";

// main function 
func main() {
	// variables => 1-(var) & 2-(:=);

	// var variableName type = value
	var name string = "Soheil";

	// variableName := value (auto type) cant use this type of variables in out of functions !!
	age := 21; // type of age int

	// const variableName type = value
	const job string = "TS And GO Dev";


	fmt.Println("my name is => " , name);
	fmt.Println("my age is => " , age);
	fmt.Println("my job is => " , job);


	// default values 
	var a string; // default value ""
	var b int; // default value 0
	var c bool; // default value false

	fmt.Println(a);
	fmt.Println(b);
	fmt.Println(c);


	// Multiple Variable Declaration 
	var num1 , num2 , num3 int = 1, 3, 5
	
	num, str := 7, "World!"

	// var (
	// 	a int
	// 	b int = 1
	// 	c string = "hello"
	// )

	fmt.Println(num1);
	fmt.Println(num2);
	fmt.Println(num3);
	fmt.Println(num);
	fmt.Println(str);
};