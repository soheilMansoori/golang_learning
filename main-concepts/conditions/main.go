package main;
import "fmt";

// main function 
func main() {
	isLogin := false;
	age := 21;

	if isLogin {
		fmt.Println("user login !!");
	} else {
		fmt.Println("use not login !!");
	};


	if(age > 18 && isLogin) {
		fmt.Println("your age ok and login")
	} else if (age > 18 && !isLogin) {
		fmt.Println("your age ok but not login");
	} else if (age > 18 || isLogin) {
		fmt.Println("ok")
	} else {
		fmt.Println("error")
	}

}