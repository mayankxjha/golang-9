package main

import "fmt"

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func closure_() {

	nextInt := intSeq()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	nextInt = intSeq()
	fmt.Println(nextInt())
}
