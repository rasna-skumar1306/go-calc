package main

import (
	"fmt"

	"github.com/go-init/calculatorGO/add"
	"github.com/go-init/calculatorGO/diff"
	"github.com/go-init/calculatorGO/sub"
)

func main() {
	s := []int{2, 3, 5, 7, 11}
	f := []float64{2.1, 3.2, 5.3, 7.4, 8.5}
	ff := []float32{2.1, 3.2, 5.3, 7.4, 8.5}
	v := add.Addi(s...)
	fv := add.Addf(f...)
	sc := diff.Diffi(s...)
	sf := diff.Difff32(ff...)
	sf64 := diff.Difff64(f...)
	si := sub.Subi(s...)
	df := sub.Subf32(ff...)
	df64 := sub.Subf(f...)
	fmt.Println(v, fv, sc, sf, sf64, si, df, df64)
	fmt.Printf("%T,%T,%T", sc, sf, sf64)
}
