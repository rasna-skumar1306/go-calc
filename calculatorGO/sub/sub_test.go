package sub

import (
	"fmt"
	"testing"
)

func TestSubi(t *testing.T) {
	n := Subi([]int{2, 3, 5, 7, 11}...)
	if n != -28 {
		t.Error("got", n, "expected", -28)
	}
}

func TestSubf(t *testing.T) {
	n := Subf([]float64{2.1, 3.2, 5.3, 7.4, 8.5}...)
	if n != -26.5 {
		t.Error("got", n, "expected", -26.5)
	}
}

func TestSubf32(t *testing.T) {
	n := Subf32([]float32{2.1, 3.2, 5.3, 7.4, 8.5}...)
	if n != -26.5 {
		t.Error("got", n, "expected", -26.5)
	}
}

func ExampleSubi() {
	fmt.Println(Subi([]int{2, 3, 5, 7, 11}...))
	//Output:
	//-28
}
func ExampleSubf() {
	fmt.Println(Subf([]float64{2.1, 3.2, 5.3, 7.4, 8.5}...))
	//Output:
	//-26.5
}
func ExampleSubf32() {
	fmt.Println(Subf32([]float32{2.1, 3.2, 5.3, 7.4, 8.5}...))
	//Output:
	//-26.5
}

func BenchmarkSubi(b *testing.B) {
	c := []int{2, 3, 5, 7, 11}
	for i := 0; i < b.N; i++ {
		Subi(c...)
	}
}
func BenchmarkSubf(b *testing.B) {
	c := []float64{2.1, 3.2, 5.3, 7.4}
	for i := 0; i < b.N; i++ {
		Subf(c...)
	}
}
func BenchmarkSubf32(b *testing.B) {
	c := []float32{2.1, 3.2, 5.3, 7.4}
	for i := 0; i < b.N; i++ {
		Subf32(c...)
	}
}
