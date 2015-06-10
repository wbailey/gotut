package main

import (
	"fmt"
	"mymath"
	"mystr"
)

func main() {
	var plink int8 = 9
	plunk := 10
	splat := 15.533

	fmt.Printf("Hello, World. Sqrt(2) = %v\n", mymath.Sqrt(2))
	fmt.Printf("plink = %s\n", plink)
	fmt.Printf("plunk = %s\n", plunk)
	fmt.Printf("splat = %f\n", splat)

	s := "string"

	s2 := "B" + s[2:]

	fmt.Println(s2)

	c := []byte(s)
	c[0] = 'S'

	s2 = string(c)

	fmt.Println(s2)

	s += "s"

	fmt.Println(s)

	s = "c" + s[0:3]

	fmt.Println(s)

	v := mystr.Replace(s, "rank", 2)

	fmt.Println(v)

	a := 1

	Add1(a)

	fmt.Println(a)

	// mymath.IsPublic()
}
