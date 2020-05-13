package diff

import (
	"fmt"
	"testing"
)

func TestDiffi(t *testing.T) {
	n := make([]int, 5)
	var er error
	var v []int
	v, er = ([]int{-1, -2, -2, -4}), nil
	n, err := Diffi([]int{2, 3, 5, 7, 11}...)
	for i, c := range n {
		if c != v[i] {
			t.Error("got", c, err, "expected", v[i], er)
		}
	}

}

func TestDifff32(t *testing.T) {
	n := make([]float32, 5)
	var v []float32
	var er error
	v, er = []float32{-1.1000001, -2.1000001, -2.1, -1.0999999}, nil
	n, err := Difff32([]float32{2.1, 3.2, 5.3, 7.4, 8.5}...)
	for i, c := range n {
		if c != v[i] {
			t.Error("got", c, "expected", v[i])
		}
	}
	if er != err {
		t.Error("got", err, "expected", er)
	}
}

func TestDifff64(t *testing.T) {
	n := make([]float64, 5)
	var v []float64
	var er error
	v, er = []float64{-1.1, -2.0999999999999996, -2.1000000000000005, -1.0999999999999996}, nil

	n, err := Difff64([]float64{2.1, 3.2, 5.3, 7.4, 8.5}...)
	if err != nil {
		return
	}
	for i, c := range n {
		if c != v[i] {
			t.Error("got", c, "expected", v[i])
		}
	}
	if er != err {
		t.Error("got", err, "expected", er)
	}
}

func ExampleDiffi() {
	fmt.Println(Diffi([]int{2, 3, 5, 7, 11}...))
	//Output:
	//[-1 -2 -2 -4]
}

func ExampleDifff32() {
	fmt.Println(Difff32([]float32{2.1, 3.2, 5.3, 7.4, 8.5}...))
	//Output:
	//[-1.1000001 -2.1000001 -2.1 -1.0999999]
}

func ExampleDifff64() {
	fmt.Println(Difff64([]float64{2.1, 3.2, 5.3, 7.4, 8.5}...))
	//Output:
	//[-1.1 -2.0999999999999996 -2.1000000000000005 -1.0999999999999996]
}
func BenchmarkDiffi(b *testing.B) {
	c := []int{2, 3, 5, 7, 11}
	for i := 0; i < b.N; i++ {
		Diffi(c...)
	}
}

func BenchmarkDifff32(b *testing.B) {
	c := []float32{2.1, 3.2, 5.3, 7.4}
	for i := 0; i < b.N; i++ {
		Difff32(c...)
	}
}

func BenchmarkDifff64(b *testing.B) {
	c := []float64{2.1, 3.2, 5.3, 7.4}
	for i := 0; i < b.N; i++ {
		Difff64(c...)
	}
}
