package main

import (
	"fmt"
	"bufio"
	"os"
)

func main()  {
	introMessage := "welcome to user input program"
	fmt.Println(introMessage)
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("enter the rating for our services")
	// now, somebody is listening or read from it(bufio)

	// now, whatever this reader actually reads, i want to store that into a variable. and this is where the comma ok/comma err syntax comes in
	input, _ := reader.ReadString('\n') 	// input, err :=   -> this is also valid
	// the argument inside ReadString means that we are gonna keep on reading till \n; it is an ender

	fmt.Println("thank you giving us rating, ", input)
	fmt.Printf("the type of rating is %T", input)  // it will be obviously string

	// in comma ok syntax, this is also valid -> _, err := reader.ReadString('\n')   // this is used when we dont care about input and care about errors

	// if we dont care about anything in go, we just use _ for that
}