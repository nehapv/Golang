//Program to find Area of Rectangle and Squarex
package main

import "fmt"

var area int

func main() {
	var l, b int
	fmt.Print("Enter Length of Rectangle : ")
	fmt.Scan(&l)
	fmt.Print("Enter Breadth of Rectangle : ")
	fmt.Scan(&b)
	area = l * b
	fmt.Println("Area of Rectangle : ", area)

	fmt.Print("Enter Length of Square : ")
	fmt.Scan(&l)
	area = l * l
	fmt.Print("Area of Square : ", area)
}
