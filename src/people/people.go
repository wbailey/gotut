package people

import "math"

type Person struct {
	name string
	age  int
}

func Older(p1 Person, p2 Person) (older Person, diff int) {
	diff = math.Abs(p1.age - p2.age)

	if p1.age > p2.age {
		older := p1
	} else {
		older := p2
	}

	return older, diff
}
