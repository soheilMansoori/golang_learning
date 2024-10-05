package main;
import "fmt";

// main function
func main() {
	// Output Functions

	// fmt.Print(value , ...value) => print value with out "\n" 
	// fmt.Println(value , ...value) => print value with "\n"
	// fmt.Printf(string , variableValueType , variableValueName) => print value with out "\n" but can use to find type or values

	const name string = "Soheil";
	
	// fmt.Printf("%T , %v" , arg_1 , arg_2) =>   
	fmt.Printf("name variable type :%T and value is => %v \n", name , name);

	isLogin := false
	fmt.Printf("isLogin variable type :%T and value is => %v \n", isLogin , isLogin);
};