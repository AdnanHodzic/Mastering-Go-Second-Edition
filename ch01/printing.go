package main

import (
	"fmt"
)

func main() {
	v1 := "123"
	v2 := 123
	v3 := "Have a nice day\n"
	v4 := "abc"

	fmt.Println("\n--- print ---\n")
	fmt.Print(v1, v2, v3, v4)
	fmt.Println("\n--- println  ---\n")
	fmt.Println(v1, v2, v3, v4)
	fmt.Println("\n--- print  ---\n")
	fmt.Print(v1, " ", v2, " ", v3, " ", v4, "\n")
	fmt.Println("\n--- printf  ---\n")
	fmt.Printf("%s%d %s %s\n", v1, v2, v3, v4)
}
