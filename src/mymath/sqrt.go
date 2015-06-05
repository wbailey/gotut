package mymath

import "fmt"

func Sqrt(x float64) float64 {
	z := 0.0

	for i := 0; i < 1000; i++ {
		z -= (z*z - x) / (2 * x)
	}

	return z
}

func IsPublic() {
	var i int8
	for i = 0; i < 10; i++ {
		fmt.Println(i)
	}

	if x := 2; x == 2 {
		fmt.Printf("this is public\n")
	}
}
