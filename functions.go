package main

import "fmt"

func plus(a int, b int) int {

	return a + b
}

func plusPlus(a, b, c int) int {
	return a + b + c
}

func functions_() {

	res := plus(1, 2)
	fmt.Println("1+2 =", res)

	res = plusPlus(1, 2, 3)
	fmt.Println("1+2+3 =", res)
}

func vals() (int, int) {
	return 3, 7
}

func functions_multi() {

	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	_, c := vals()
	fmt.Println(c)
}

func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0

	for itr, num := range nums {
		total += num
		fmt.Println("summed iteration: ", itr)
	}
	fmt.Println(total)
}

func function_variadic() {

	sum(1, 2)
	sum(1, 2, 3)

	nums := []int{1, 2, 3, 4}
	sum(nums...)
}
