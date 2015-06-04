package main

import (
	"fmt"
	"mymath"
)

func main() {
	fmt.Printf("Hello, World. Sqrt(2) = %v\n", mymath.Sqrt(2))

	mymath.IsPublic()

	mymath.Notpublic()
}
