package main

import (
	"fmt"

	"github.com/go-init/calculatorGO/add"
	"github.com/go-init/calculatorGO/diff"
)

func main() {
	s := []int{1, 2, 3123, 124, 4, 5, 45, 56, 6, 67, 56, 3, 2}
	f := []float64{1.5435, 345.76, 345345.7664, 345345.767}
	v := add.Addi(s...)
	fv := add.Addf(f...)
	sc := diff.Diffi(s...)
	fmt.Println(v, fv, sc)
}
