package mymath

func Add1(a *int) int {
	*a = *a + 1
	return *a
}
