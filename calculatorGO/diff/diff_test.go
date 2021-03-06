package diff

import (
	"fmt"
	"testing"
)

func TestDiffi(t *testing.T) {
	n := make([]int, 5)
	var v []int
	v = ([]int{-1, -2, -2, -4})
	n, err := Diffi([]int{2, 3, 5, 7, 11}...)
	if err != nil {
		t.Error("got", err, "expected", nil)
	}
	for i, c := range n {
		if c != v[i] {
			t.Error("got", c, "expected", v[i])
		}
	}

}

func TestDifff32(t *testing.T) {
	n := make([]float32, 5)
	var v []float32
	v = []float32{-1.1000001, -2.1000001, -2.1, -1.0999999}
	n, err := Difff32([]float32{2.1, 3.2, 5.3, 7.4, 8.5}...)
	if err != nil {
		t.Error("got", err, "expected", nil)
	}
	for i, c := range n {
		if c != v[i] {
			t.Error("got", c, "expected", v[i])
		}
	}

}

func TestDifff64(t *testing.T) {
	n := make([]float64, 5)
	var v []float64
	v = []float64{-1.1, -2.0999999999999996, -2.1000000000000005, -1.0999999999999996}
	n, err := Difff64([]float64{2.1, 3.2, 5.3, 7.4, 8.5}...)
	if err != nil {
		t.Error("got", err, "expected", nil)
	}
	for i, c := range n {
		if c != v[i] {
			t.Error("got", c, "expected", v[i])
		}
	}

}

func ExampleDiffi() {
	d, _ := Diffi([]int{2, 3, 5, 7, 11}...)
	fmt.Println(d)
	//Output:
	//[-1 -2 -2 -4]
}

func ExampleDifff32() {
	d, _ := Difff32([]float32{2.1, 3.2, 5.3, 7.4, 8.5}...)
	fmt.Println(d)
	//Output:
	//[-1.1000001 -2.1000001 -2.1 -1.0999999]
}

func ExampleDifff64() {
	d, _ := Difff64([]float64{2.1, 3.2, 5.3, 7.4, 8.5}...)
	fmt.Println(d)
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
