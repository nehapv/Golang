// Go program
// how to concatenate strings
package main

import "fmt"

func main() {

	var str1 string
	str1 = "Welcome!"

	var str2 string
	str2 = "GeeksforGeeks"

	fmt.Println("New string 1: ", str1+str2)

	str3 := "Geeks"
	str4 := "Geeks"

	result := str3 + "for" + str4

	fmt.Println("New string 2: ", result)

}
