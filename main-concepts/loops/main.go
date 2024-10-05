package main;
import "fmt";

// main function
func main() {
	// for statement1; statement2; statement3 {
	// 	 code to be executed for each iteration
	// }

	// simple loop 
	for i :=0 ; i <= 5; i++ {
		fmt.Println("i => " , i);
	};
	
	// find even numbers in 10 use(continue)
	for i := 0; i <= 10; i++ {
		// if num odd continue
		if i % 2 != 0 {
			continue;
		};

		fmt.Println("event num is => " , i);
	};

	// break in loops
	for num := 0; num < 5; num++ {
		// break loop
		if num == 3 {
		  break;
		};
	   fmt.Println("num" ,num);
	}; 
	
	// create array 
	var languages = [] string {"TS" , "JS" , "Kotlin" , "PHP" , "Python" , "Java" , "Go" , "C++"};

	// get item of array way-1 (for)
	for i := 0; i < len(languages); i++ {
		fmt.Println("item => " , languages[i]);
	};

	// get item of array way-2 (range)
	for index, value := range languages {
		result := fmt.Sprintf("value : %s and index : %d" , value , index);	
		fmt.Println("result => " , result);
	};
};



