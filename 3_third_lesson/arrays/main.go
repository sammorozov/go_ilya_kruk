package main

import "fmt"

func main() {
	var a [3]int

	a[0] = 1
	a[1] = 2
	a[2] = 200

	modifyMe(&a)

	fmt.Println(a)

	for index, el := range a {
		fmt.Println(index, el)
	}

	r := [...]rune{9: 'a'}
	fmt.Println(r[9], r[8], r[0], r[9])

	s1 := r[3:10]
	fmt.Println(s1[6], len(s1))

}

func modifyMe(arr *[3]int) {
	arr[1] = 100

}
