package main

import (
	"fmt"
)

func main()  {
	fmt.Println("welcome to pointers lesson")
	// var ptr *int   // we can even do: var ptr *string
	// fmt.Println("the value of pointer is: ", ptr)   
	// the default value is 'the value of pointer is:  <nil>'

	// in above, we created a pointer that doesnt point at anything(points at no memory address) and thats why its value is nil
	// now lets create a pointer that points at memory address of a variable
	myNumber := 23
	var ptr = &myNumber
	fmt.Println("the value of pointer is: ", ptr)   //the value of pointer is:  0xc00008c0a8
	fmt.Println("the value of pointer is: ", *ptr)  //the value of pointer is:  23

	*ptr = *ptr + 2
	fmt.Println("the value of myNumber is: ", myNumber)  //the value of myNumber is:  25
}