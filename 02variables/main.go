package main

import "fmt"

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

}
