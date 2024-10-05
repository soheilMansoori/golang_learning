package main;
import "fmt";

// send message 
func sendMessage(message string) string {
	return message;
};

// add function
func add(num1 int , num2 int) int {
	return num1 + num2;
}


func main() {
	// functions in go
	// func FunctionName(param1 type, param2 type) type {
	// 	// code to be executed
	// 	return output
	// } 

	fmt.Println("message is => " , sendMessage("Hello World"));
	fmt.Println("add result => " , add(1,3));
};