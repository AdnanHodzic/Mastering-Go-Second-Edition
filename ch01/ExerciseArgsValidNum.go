// find the sum of all command-line arguments that are valid numbers.
package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func main() {
	argsWithoutApp := os.Args[1:]
	if len(os.Args) < 3 {
		fmt.Println("Please provide 2 numbers.")
		os.Exit(1)
	} else if len(os.Args) >= 11 {
		fmt.Println("You've entered more then 10 arguments.")
		os.Exit(1)
	}

	arguments := os.Args
	var err error = errors.New("An error")
	k := 1
	var n float64
	for err != nil {
		if k >= len(arguments) {
			fmt.Printf("None of the arguments is a number! Need: %T\n", n)
			return
		}
		n, err = strconv.ParseFloat(arguments[k], 64)
		k++
	}

	sum := 0
	for i := 1; i < len(arguments); i++ {
		temp, _ := strconv.Atoi(arguments[i])
		sum = sum + temp
	}

	fmt.Println("Total sum of provided", argsWithoutApp, "numbers is:", sum)
}
