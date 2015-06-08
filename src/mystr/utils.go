package mystr

import "fmt"

func Replace(src string, rpl string, count int) string {
	convert := []byte(src)

	fmt.Println(len(convert))

	for i := 0; i < len(convert); i++ {
		fmt.Println(convert[i])
	}

	fin := string(convert)

	return fin
}
