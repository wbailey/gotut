package mymath

func Add1(a *int) int {
	*a += 1
	return *a
}

type testInt func(int) bool

func IsOdd(a int) bool {
	return a%2 == 1
}

func IsEven(a int) bool {
	return !IsOdd(a)
}

func Filter(vals []int, f testInt) []int {
	var result []int

	for _, value := range vals {
		if f(value) {
			result = append(result, value)
		}
	}

	return result
}
