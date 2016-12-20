package main

import (
	"fmt"
)

type int_a struct {
	a int
	b int
	c int
}

func main() {
	struct_int := int_a{1, 2, 3}
	fmt.Printf("%d\n", struct_int.a)
}
