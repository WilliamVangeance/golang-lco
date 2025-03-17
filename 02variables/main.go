package main

import "fmt"

const LoginToken = "somegarbagevalue" // global variable

// notallowed := 30000	// use of this operator := is not allowed outside of function or method

func main() {
	var isequal uint8 = 255 // unsigned integer, containes 0-255 , similiarily we have uint16 , uint32 etc.,
	fmt.Println(isequal)
	fmt.Printf("vairable is of type %T \n", isequal)

	var username string = "yugal"
	fmt.Println(username)
	fmt.Printf("variable username is of type %T \n", username)

	var anothervariable int
	fmt.Println(anothervariable) // instead of any garbage value like in other programming languages it returns zero

	var anotherstring string
	fmt.Println(anotherstring) // it returns blank space for strings

	var randomname = "teaovercoffee" // implicit type , here we're not giving the variable type
	fmt.Println(randomname)
	fmt.Printf("variable is of type %T \n", randomname)

	var randomfloat = 3.1389
	fmt.Println(randomfloat)
	fmt.Printf("variable is of type %T \n", randomfloat)

	withoutvar := "variable without var" // no var type declaration
	fmt.Println(withoutvar)

	novarint := 300000 // inside any method we're allowed to use this short variable declaration operator
	fmt.Println(novarint)

	fmt.Println(LoginToken) // printing the constant value

}
