package main

import "fmt"

type base struct {
	num int
	str string
}

func (b base) describe() string {
	return fmt.Sprintf("base with num=%v and with str=%v", b.num, b.str)
}

type container struct {
	base
	str string
}

func struct_embed() {

	co := container{
		base: base{
			num: 1,
			str: "wowo",
		},
		str: "some name",
	}

	fmt.Printf("co={base.num: %v, base.str: %v, str: %v}\n", co.base.num, co.base.str, co.str)

	fmt.Println("also num:", co.base.num)

	fmt.Println("describe:", co.describe())

	type describer interface {
		describe() string
	}

	var d describer = co
	fmt.Println("describer:", d.describe())
}
