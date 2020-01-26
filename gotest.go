// gotest.go
// main file for testing features of Go language.

package main

// Imported packages.
import (
	"fmt"
)

// Package level variables.
var (
	message string = "The message"
)

// Initialization function.
func init() {
	fmt.Println("Initialization Started...")

	fmt.Println("Initialization Complete.")
}

// Main function - Entrypoint of application.
func main() {
	fmt.Println("Application Started...\n")

	i := 42
	j := float64(i)

	fmt.Printf("i: %v, %T\n", i, i)
	fmt.Printf("j: %v, %T\n", j, j)

	myArr := [3]int{}
	myArr2 := [...]int{2,4,6}
	myArr[0] = 12
	myArr[2] = 3
	myArr[1] = 9

	fmt.Printf("Array of %v indices: %v, %v, %v\n", len(myArr), myArr[0], myArr[1], myArr[2])
	fmt.Println(myArr)
	fmt.Println(myArr2)

	mySlice :=[]int{3,6,9}
	mySlice = append(mySlice, 12)
	fmt.Printf("Slice: %v\n", mySlice)



//	displayMsg()
	fmt.Println("\nApplication Terminated.\n")
}
