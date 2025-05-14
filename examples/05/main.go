package main

// go to terminal and 'go env'
// look for GOOS value(not GOHOSTOS) for info about currently used operating system

/*
// to build executables(exe for windows, etc) of our program, we have to do the following on terminal
// go build

// if we need to build executables for a specific operating system and not the current one
// GOOS="windows" go build

// for mac, GOOS="darwin"
// for linux, GOOS="linux"

now lets take one of the previous program as example to build exe for windows
*/

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
	newRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("added 1 to your rating: ", newRating + 1)
	}
}