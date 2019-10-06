// keep reading (int) until it gets the word END as input
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Enter stuff ...")
	for scanner.Scan() {
		fmt.Println(">", scanner.Text())
		notNumber, err := strconv.Atoi(scanner.Text())
		if scanner.Text() == "END" {
			fmt.Println("Reched the END, goodbye")
			os.Exit(1)
		} else if err != nil {
			fmt.Println("Value is not an integer:", notNumber, "\ntry agian ...")

		}
	}
}
