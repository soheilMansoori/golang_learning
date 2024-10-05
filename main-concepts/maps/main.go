package main;
import "fmt";

func main() {
	var soheilUser = map [string] string {
		"name" : "soheil",
		"age" : "21",
		"job" : "TS and Go Dev",
	};

	// update data 
	soheilUser["name"] = "soheil user updated";

	// delete item 
	delete(soheilUser , "job");

	fmt.Println("soheil user => " , soheilUser);
	fmt.Println("soheil name => " , soheilUser["name"]); // error if soheilUser.name !!

	// create item
	soheilUser["job"] = "TS and Go Dev";	

	// use loops in maps 
	for key, value := range soheilUser {
		fmt.Printf("%v : %v, \n", key, value);
	};
}