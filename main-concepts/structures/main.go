package main;
import "fmt";

// struct (class)
type Users struct {
	name string
	age int
	job string
}


func main() {
	var soheilUser Users;

	fmt.Println("soheilUser => " , soheilUser); // "" 0 ""
	
	// fill data
	soheilUser.name = "soheil";
	soheilUser.age = 21;
	soheilUser.job = "TS and Go Dev";

	fmt.Println("soheilUser => " , soheilUser);

}