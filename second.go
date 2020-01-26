// second.go
// test file 
package main

import (
	"fmt"
)

func displayMsg() {
	var msg string
	msg = message
//	msg = "theMessage"

	fmt.Printf("Message: %v, %T\n", msg, msg)
}
