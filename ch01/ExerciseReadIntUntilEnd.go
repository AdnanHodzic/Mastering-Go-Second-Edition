// keep reading (int) until it gets the word END as input
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Enter stuff ...")
	for scanner.Scan() {
		fmt.Println(">", scanner.Text())
		if scanner.Text() == "END" {
			fmt.Println("Reched the END, goodbye")
			os.Exit(1)
		}
	}
}
