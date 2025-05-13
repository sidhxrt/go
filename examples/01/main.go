package main

import "fmt"

// variableA := 123    // this will throw error as walrus operators are allowed only inside the method
// var variableB = 123     // this is allowed
// var variableC int = 123   // this is allowed

// once assigned, we cannot change the value of const. var variables can be changed during the period of execution of program
const LoginToken = "abcdef"
// the reason why we made L caps is because this is now a public variable
// const LoginToken = "abc" is equivalent to 'public const logintoken = "abc" of java'


func main()  {
	var username string = "sid"
	fmt.Printf("the variable is of the type: %T \n", username)

	var isUsername bool = true
	fmt.Println(isUsername)
	fmt.Printf("the variable is of the type: %T \n", isUsername)

	var smallVal uint8 = 255     // uint is unsigned integer
	fmt.Println(smallVal)
	fmt.Printf("the variable is of the type: %T \n", smallVal)

	var smallFloat float32 = 232.56234891623326485
	fmt.Println(smallFloat)
	fmt.Printf("the variable is of the type: %T \n", smallFloat)

	var largeFloat float64 = 232.56234891623326485
	fmt.Println(largeFloat)
	fmt.Printf("the variable is of the type: %T \n", largeFloat)

	// default values and some aliases
	var a int         // default value is 0
	fmt.Println(a)
	fmt.Printf("type: %T \n", a)

	var b string
	fmt.Println(b)
	fmt.Printf("type: %T \n", b)

	// there is an implicit type or implicit way of declaring variables
	// implicit type
	var c = "wassup"
	fmt.Println(c)
	fmt.Printf("type: %T \n", c)

	// no var type
	numberOfUsers := 3000.00
	fmt.Println(numberOfUsers)

	fmt.Println(LoginToken)
	fmt.Printf("type: %T \n", LoginToken)

}