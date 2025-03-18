package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "welcome to user input program"
	fmt.Println(welcome)
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("enter the rating of our pizza :")
	input, _ := reader.ReadString('\n')
	fmt.Println("thanks for rating , ", input)
	// comma ok syntax / comma error syntax  input, err   / input, _  / _, err
}
