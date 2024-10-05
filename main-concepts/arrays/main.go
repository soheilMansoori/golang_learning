package main;
import "fmt";

// main function 
func main() {
	/*
		var arrayName = [length]datatype{values} // here length is defined
		or
		var arrayName = [...]datatype{values} // here length is inferred
	*/

	// array of numbers 
	var numbers = [3] int {1,2,3};
	fmt.Println("numbers => " , numbers);


	// array of strings
	languages := [5] string {"TS" , "JS" , "GO" , "PYTHON" , "PHP"};
	fmt.Println("languages => " , languages);

	// array of default values
	defaultArray := [5] int {}; // [0,0,0,0,0]
	fmt.Println("defaultArray => " , defaultArray);


	// array with out default length (empty array)
	var emptyArray = [...] int {};
	fmt.Println("emptyArray => " , emptyArray);
	
	// get item of array with index
	fmt.Println("first item of numbers languages => " , languages[0]);

	// change item of array with index 
	fmt.Println("last item of languages before change => " , languages[4]);
	languages[4] = "KOTLIN";
	fmt.Println("last item of languages after change => " , languages[4]);


	// create array with custom index
	var carBrands = [5] string {1:"BMV" , 4 : "BENZ"} // ["" "BMV" "" "" "BENZ"]
	fmt.Println("carBrands => " , carBrands);


	// get array length
	fmt.Println("carBrands length => " , len(carBrands));
}