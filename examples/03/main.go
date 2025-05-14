package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main()  {
	msg1 := "welcome to pizza place"
	msg2 := "please give us a rating from 1 to 5"
	fmt.Println(msg1)
	fmt.Println(msg2)
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	fmt.Println("thanks for rating us with,", input)
	/*
	newRating, err := strconv.ParseFloat(input, 64)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("added 1 to your rating: ", newRating + 1)
	}

	the above code will throw the following error: strconv.ParseFloat: parsing "4\r\n": invalid syntax
	we need to remove the \n and stuff, for this we will use another go package(this one is very powerful and useful) called 'strings'
	*/
	newRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("added 1 to your rating: ", newRating + 1)
	}
}