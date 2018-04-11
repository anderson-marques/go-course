package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// fmain short cut
func main() {
	//	var s string
	// Pasing by reference
	//	fmt.Scanln(&s) // It gets only the first token delimited by space
	// Printing by value
	//	fmt.Println(s)

	// Declare the Reader from Standard Input
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter your name: ")
	str, _ := reader.ReadString('\n')
	fmt.Println(str)

	fmt.Print("Enter your age: ")
	age, _ := reader.ReadString('\n')
	f, err := strconv.ParseFloat(strings.TrimSpace(age), 64)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(f)
	}

}
