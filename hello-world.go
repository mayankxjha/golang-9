package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("hello world")
	fmt.Println(strings.Repeat("-", 50))
	values()
	fmt.Println(strings.Repeat("-", 50))
	variables()
	fmt.Println(strings.Repeat("-", 50))
	constants()
	fmt.Println(strings.Repeat("-", 50))
	for_loop()
	fmt.Println(strings.Repeat("-", 50))
	if_else()
	fmt.Println(strings.Repeat("-", 50))
	switch_()
	fmt.Println(strings.Repeat("-", 50))
	arrays()
	fmt.Println(strings.Repeat("-", 50))
	slices_()
	fmt.Println(strings.Repeat("-", 50))
	maps_()
	fmt.Println(strings.Repeat("-", 50))
	functions_()
	fmt.Println(strings.Repeat("-", 50))
	functions_multi()
	function_variadic()
	fmt.Println(strings.Repeat("-", 50))
}
